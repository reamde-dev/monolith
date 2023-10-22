package monolith

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"sync/atomic"

	"reamde.dev/monolith/internal/xsync"
)

type RPC struct {
	writerQueue *xsync.Queue[*pendingWrite]
	readerQueue *xsync.Queue[*pendingRead]

	requests   *xsync.Map[uint64, *pendingRequest]
	requestSeq uint64

	close     chan struct{}
	closeDone chan struct{}
}

type pendingRequest struct {
	dst  []byte
	err  error
	done chan struct{}
}

type pendingWrite struct {
	seq  uint64
	path string
	buf  []byte
	wait bool
	err  error
	done chan struct{}
}

type pendingRead struct {
	seq  uint64
	path string
	body []byte
}

func NewRPC(
	sess *Session,
) *RPC {
	rpc := &RPC{
		writerQueue: xsync.NewQueue[*pendingWrite](10),
		readerQueue: xsync.NewQueue[*pendingRead](10),
		requests:    xsync.NewMap[uint64, *pendingRequest](),
		close:       make(chan struct{}),
		closeDone:   make(chan struct{}),
	}
	rpc.handle(sess)
	return rpc
}

func (c *RPC) handle(sess *Session) {
	writerDone := make(chan error)
	readerDone := make(chan error)

	go func() {
		writerDone <- c.writeLoop(sess)
		close(writerDone)
	}()

	go func() {
		readerDone <- c.readLoop(sess)
		close(readerDone)
	}()

	go func() {
		// Note: do NOT close the session, it is owned by the caller
		// wait for the close signal
		<-c.close
		// close reader and writer
		c.writerQueue.Close()
		c.readerQueue.Close()
		// wait for reader and writer to finish
		<-writerDone
		// TODO: reader loop is currently not closing
		// <-readerDone
		// wrap everything up
		close(c.close)
		close(c.closeDone)
	}()
}

func (c *RPC) Request(ctx context.Context, path string, payload []byte) ([]byte, error) {
	pr := &pendingRequest{
		done: make(chan struct{}),
	}

	requestSeq := c.next()
	c.requests.Store(requestSeq, pr)

	err := c.write(requestSeq, path, payload, false)
	if err != nil {
		close(pr.done)
		c.requests.Delete(requestSeq)
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-pr.done:
	}

	return pr.dst, pr.err
}

func (c *RPC) next() uint64 {
	return atomic.AddUint64(&c.requestSeq, 1)
}

func (c *RPC) write(seq uint64, path string, payload []byte, wait bool) error {
	pw := &pendingWrite{
		seq:  seq,
		path: path,
		buf:  payload,
		wait: wait,
		done: make(chan struct{}),
	}

	err := c.writerQueue.Push(pw)
	if err != nil {
		if errors.Is(err, xsync.ErrQueueClosed) {
			return fmt.Errorf("writer queue closed: %w", io.EOF)
		}
		return fmt.Errorf("writer queue error: %w", err)
	}

	if !wait {
		return nil
	}

	<-pw.done
	return pw.err
}

func (c *RPC) writeLoop(conn *Session) error {
	for {
		pw, err := c.writerQueue.Pop()
		if err != nil {
			if errors.Is(err, xsync.ErrQueueClosed) {
				return nil
			}
			return fmt.Errorf("pop from writer queue: %w", err)
		}

		// TODO: probably too expensive
		body := []byte{}
		body = binary.AppendUvarint(body, pw.seq)
		body = binary.AppendUvarint(body, uint64(len(pw.path)))
		body = append(body, pw.path...)
		body = binary.AppendUvarint(body, uint64(len(pw.buf)))
		body = append(body, pw.buf...)
		_, err = conn.write(body)
		if pw.wait {
			pw.err = err
			close(pw.done)
		}
		if err != nil {
			return fmt.Errorf("write error: %w", err)
		}
	}
}

func (c *RPC) readNextUInt(r *bytes.Reader) (uint64, error) {
	seq, err := binary.ReadUvarint(r)
	if err != nil {
		return 0, fmt.Errorf("read seq error: %w", err)
	}

	return seq, nil
}

func (c *RPC) readNextBytes(r *bytes.Reader) ([]byte, error) {
	size, err := binary.ReadUvarint(r)
	if err != nil {
		return nil, fmt.Errorf("read size error: %w", err)
	}

	if size == 0 {
		return nil, nil
	}

	data := make([]byte, size)
	n, err := io.ReadFull(r, data)
	if err != nil {
		return nil, fmt.Errorf("read data error: %w", err)
	}

	if n != len(data) {
		return nil, fmt.Errorf("read data error: %w", io.ErrUnexpectedEOF)
	}

	return data, nil
}

func (c *RPC) readLoop(conn *Session) error {
	for {
		message, err := conn.read()
		if err != nil {
			return fmt.Errorf("read error: %w", err)
		}

		reader := bytes.NewReader(message)

		seq, err := c.readNextUInt(reader)
		if err != nil {
			return fmt.Errorf("read seq error: %w", err)
		}

		path, err := c.readNextBytes(reader)
		if err != nil {
			return fmt.Errorf("read path error: %w", err)
		}

		data, err := c.readNextBytes(reader)
		if err != nil {
			return fmt.Errorf("read data error: %w", err)
		}

		pr, exists := c.requests.Load(seq)
		if exists {
			c.requests.Delete(seq)
		}

		if seq == 0 || !exists {
			err := c.readerQueue.Push(&pendingRead{
				seq:  seq,
				path: string(path),
				body: data,
			})
			if err != nil {
				if errors.Is(err, xsync.ErrQueueClosed) {
					return nil
				}
				return fmt.Errorf("handler encountered an error: %w", err)
			}
			continue
		}

		pr.dst = data
		close(pr.done)
	}
}

func (c *RPC) Read() (path string, body []byte, callback func([]byte) error, err error) {
	pr, err := c.readerQueue.Pop()
	if err != nil {
		if errors.Is(err, xsync.ErrQueueClosed) {
			return "", nil, nil, fmt.Errorf("reader queue closed: %w", io.EOF)
		}
		return "", nil, nil, fmt.Errorf("reader queue error: %w", err)
	}
	return pr.path, pr.body, func(payload []byte) error {
		if pr.seq == 0 {
			return fmt.Errorf("not a request")
		}
		return c.write(pr.seq, path, payload, true)
	}, nil
}

func (c *RPC) Close() error {
	// TODO implement error?
	c.close <- struct{}{}
	<-c.closeDone
	return nil
}

package monolith

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRPC_E2E(t *testing.T) {
	// construct new mock connection between two nodes
	mc := NewMockSession(t)

	// construct a new connection for the "server"
	srv := NewRPC(mc.Server)

	// add a handler for the "server"
	go func() {
		path, msg, cb, err := srv.Read()
		require.NoError(t, err)
		require.Equal(t, "ping", path)
		require.Equal(t, "PING!", string(msg))

		err = cb([]byte("pong"))
		require.NoError(t, err)
	}()

	// construct a new connection for the "client"
	cln := NewRPC(mc.Client)

	// client writes to server
	res, err := cln.Request(context.Background(), "ping", []byte("PING!"))
	require.NoError(t, err)
	require.Equal(t, "pong", string(res))

	t.Run("closed connection returns io.EOF", func(t *testing.T) {
		// close the connections
		srv.Close()
		cln.Close()

		// client writes to server errors
		_, err := cln.Request(context.Background(), "ping", []byte("PING!"))
		require.ErrorIs(t, err, io.EOF)

		// server writes to client errors
		_, err = srv.Request(context.Background(), "ping", []byte("PING!"))
		require.ErrorIs(t, err, io.EOF)
	})
}

func TestRPC_E2E_LongMessage(t *testing.T) {
	// construct new mock connection between two nodes
	mc := NewMockSession(t)

	// construct a new connection for the "server"
	srv := NewRPC(mc.Server)

	// create a long message, longer than the buffer size
	body := make([]byte, 4096+100)
	for i := range body {
		body[i] = 'a'
	}

	// add a handler for the "server"
	go func() {
		path, msg, cb, err := srv.Read()
		require.NoError(t, err)
		assert.Equal(t, body, msg)
		assert.Equal(t, "ping", path)
		cb([]byte("ok"))
	}()

	// construct a new connection for the "client"
	cln := NewRPC(mc.Client)

	// client writes to server
	res, err := cln.Request(context.Background(), "ping", body)
	require.NoError(t, err)
	require.Equal(t, "ok", string(res))
}
package rpc

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"sync"
)

type (
	Request struct {
		Topic string
		Body  []byte
	}
	Response struct {
		Body  []byte
		Error string
	}
)

type Client struct {
	sync.Mutex
	conns map[string]net.Conn
}

func (c *Client) getConnection(serverAddress string) (net.Conn, error) {
	c.Lock()
	defer c.Unlock()

	if c.conns == nil {
		c.conns = make(map[string]net.Conn)
	}

	// If the connection doesn't exist or is closed, create a new one
	if conn, ok := c.conns[serverAddress]; ok && conn != nil {
		return conn, nil
	}

	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		return nil, err
	}
	c.conns[serverAddress] = conn

	return conn, nil
}

func (c *Client) closeConnection(serverAddress string) {
	c.Lock()
	defer c.Unlock()

	if conn, ok := c.conns[serverAddress]; ok && conn != nil {
		conn.Close()
		delete(c.conns, serverAddress)
	}
}

func (c *Client) SendRequest(serverAddress string, request Request) (Response, error) {
	conn, err := c.getConnection(serverAddress)
	if err != nil {
		return Response{}, fmt.Errorf("error getting connection: %w", err)
	}

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	if err := encoder.Encode(&request); err != nil {
		c.closeConnection(serverAddress) // Close connection in case of error
		return Response{}, fmt.Errorf("error encoding request: %w", err)
	}

	var response Response
	if err := decoder.Decode(&response); err != nil {
		c.closeConnection(serverAddress) // Close connection in case of error
		return Response{}, fmt.Errorf("error decoding response: %w", err)
	}

	return response, nil
}

type Server struct {
	Address string
	Handler func(request Request) Response
}

func (s *Server) handleConnection(conn net.Conn) {
	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	for {
		var request Request
		if err := decoder.Decode(&request); err != nil {
			if err != io.EOF {
				fmt.Println("Error decoding:", err)
			}
			break // exit the loop if there's a decoding error or client closed the connection
		}

		response := s.Handler(request)
		if err := encoder.Encode(&response); err != nil {
			fmt.Println("Error encoding:", err)
			break // exit the loop if there's an encoding error
		}
	}

	conn.Close()
}

func (s *Server) Listen() {
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

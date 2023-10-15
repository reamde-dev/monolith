package rpc

import (
	"encoding/gob"
	"fmt"
	"net"
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
	ServerAddress string
}

func (c *Client) SendRequest(request Request) (Response, error) {
	conn, err := net.Dial("tcp", c.ServerAddress)
	if err != nil {
		return Response{}, err
	}
	defer conn.Close()

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	if err := encoder.Encode(&request); err != nil {
		return Response{}, err
	}

	var response Response
	if err := decoder.Decode(&response); err != nil {
		return Response{}, err
	}

	return response, nil
}

type Server struct {
	Address string
	Handler func(request Request) Response
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	var request Request
	if err := decoder.Decode(&request); err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	response := s.Handler(request)
	if err := encoder.Encode(&response); err != nil {
		fmt.Println("Error encoding:", err)
	}
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

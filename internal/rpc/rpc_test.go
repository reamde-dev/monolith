package rpc

import (
	"fmt"
	"testing"
	"time"
)

func TestClientServer(t *testing.T) {
	// Start multiple servers
	serverAddresses := []string{":8081", ":8082", ":8083"}
	for _, address := range serverAddresses {
		server := &Server{
			Address: address,
			Handler: func(request Request) Response {
				return Response{
					Body: []byte(fmt.Sprintf("Topic: %s, Payload: %s", request.Topic, request.Body)),
				}
			},
		}
		go server.Listen() // start each server in a goroutine
	}

	// Allow the servers a moment to start up
	time.Sleep(1 * time.Second)

	client := &Client{}

	// We will have one test for each topic for each server
	tests := []struct {
		serverAddress string
		request       Request
		expected      Response
	}{
		{
			serverAddress: "localhost:8081",
			request: Request{
				Topic: "topic1",
				Body:  []byte("Payload for topic1"),
			},
			expected: Response{
				Body: []byte("Topic: topic1, Payload: Payload for topic1"),
			},
		},
		{
			serverAddress: "localhost:8081", // Same server to test multiple requests
			request: Request{
				Topic: "topic2",
				Body:  []byte("Payload for topic2"),
			},
			expected: Response{
				Body: []byte("Topic: topic2, Payload: Payload for topic2"),
			},
		},
		{
			serverAddress: "localhost:8082",
			request: Request{
				Topic: "topic1",
				Body:  []byte("Payload for topic1"),
			},
			expected: Response{
				Body: []byte("Topic: topic1, Payload: Payload for topic1"),
			},
		},
		{
			serverAddress: "localhost:8083",
			request: Request{
				Topic: "topic1",
				Body:  []byte("Payload for topic1"),
			},
			expected: Response{
				Body: []byte("Topic: topic1, Payload: Payload for topic1"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s-%s", tt.serverAddress, tt.request.Topic), func(t *testing.T) {
			response, err := client.SendRequest(tt.serverAddress, tt.request)
			if err != nil {
				t.Fatalf("Expected no error, but got: %v", err)
			}

			if string(response.Body) != string(tt.expected.Body) {
				t.Errorf("Expected response payload %s, but got: %s", tt.expected.Body, response.Body)
			}
		})
	}
}

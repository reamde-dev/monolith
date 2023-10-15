package rpc

import (
	"fmt"
	"testing"
	"time"
)

func TestClientServer(t *testing.T) {
	server := &Server{
		Address: ":8081",
		Handler: func(request Request) Response {
			return Response{
				Body: []byte(fmt.Sprintf("Topic: %s, Payload: %s", request.Topic, request.Body)),
			}
		},
	}
	go server.Listen() // start server in a goroutine

	// Allow the server a moment to start up
	time.Sleep(1 * time.Second)

	client := &Client{
		ServerAddress: "localhost:8081",
	}

	tests := []struct {
		request  Request
		expected Response
	}{
		{
			request: Request{
				Topic: "topic1",
				Body:  []byte("Payload for topic1"),
			},
			expected: Response{
				Body: []byte("Topic: topic1, Payload: Payload for topic1"),
			},
		},
		{
			request: Request{
				Topic: "topic2",
				Body:  []byte("Payload for topic2"),
			},
			expected: Response{
				Body: []byte("Topic: topic2, Payload: Payload for topic2"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.request.Topic, func(t *testing.T) {
			response, err := client.SendRequest(tt.request)
			if err != nil {
				t.Fatalf("Expected no error, but got: %v", err)
			}

			if string(response.Body) != string(tt.expected.Body) {
				t.Errorf("Expected response payload %s, but got: %s", tt.expected.Body, response.Body)
			}
		})
	}
}

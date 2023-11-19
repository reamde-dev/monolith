package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/proto"

	"reamde.dev/monolith/proto/example"
	"reamde.dev/monolith/proto/monolith"
)

func main() {
	// Read base64 string (assuming passed as the first argument)
	base64Message := os.Args[1]

	// Decode from base64
	bytes, err := base64.StdEncoding.DecodeString(base64Message)
	if err != nil {
		log.Fatalf("Failed to decode base64: %v", err)
	}

	// Unmarshal into Message
	var message monolith.Message
	if err := proto.Unmarshal(bytes, &message); err != nil {
		log.Fatalf("Failed to unmarshal message: %v", err)
	}

	// Unmarshal Any payload
	payload, err := message.Payload.UnmarshalNew()
	if err != nil {
		log.Fatalf("Failed to unmarshal payload: %v", err)
	}

	// Print the payload
	typeA, ok := payload.(*example.TypeA)

	// Print the type of the payload
	fmt.Printf("Payload type: %T\n", payload)
	fmt.Printf("Payload is TypeA: %v\n", ok)
	fmt.Printf("Payload: %v\n", typeA)
}

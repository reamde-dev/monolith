package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/anypb"

	"reamde.dev/monolith/proto/example"
	"reamde.dev/monolith/proto/monolith"
)

func main() {
	// Create an instance of TypeA
	typeA := &example.TypeA{Data: "Hello, World!"}

	// Marshal TypeA into an Any type
	anyTypeA, err := anypb.New(typeA)
	if err != nil {
		log.Fatalf("Failed to marshal TypeA: %v", err)
	}

	// Store TypeA in the Message payload
	message := &monolith.Message{
		Payload: anyTypeA,
	}

	// ... Similarly for TypeB ...

	// Example of retrieving the payload
	payload, err := message.Payload.UnmarshalNew()
	if err != nil {
		log.Fatalf("Failed to unmarshal payload: %v", err)
	}

	switch p := payload.(type) {
	case *example.TypeA:
		fmt.Printf("Payload is TypeA: %v\n", p)
	case *example.TypeB:
		fmt.Printf("Payload is TypeB: %v\n", p)
	default:
		fmt.Printf("Unknown type: %T\n", p)
	}
}

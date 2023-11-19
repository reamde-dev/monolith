package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"

	"reamde.dev/monolith/proto/example"
	"reamde.dev/monolith/proto/monolith"
)

func main() {
	// Create an instance of TypeA
	typeA := &example.TypeA{Data: "Example data"}

	// Marshal TypeA into an Any type
	anyTypeA, err := NewAny(typeA)
	if err != nil {
		log.Fatalf("Failed to marshal TypeA: %v", err)
	}

	// Create Message with TypeA as payload
	message := &monolith.Message{
		Payload: anyTypeA,
	}

	// Marshal Message to protobuf binary format
	bytes, err := proto.Marshal(message)
	if err != nil {
		log.Fatalf("Failed to marshal message: %v", err)
	}

	// Convert to base64 for easy text transmission
	base64Message := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println(base64Message)
}

// Any is a wrapper around anypb.New that replaces the default type URL with
// a monolith specific one.
// https://developers.google.com/protocol-buffers/docs/proto3#any
func NewAny(msg protoreflect.ProtoMessage) (*anypb.Any, error) {
	anyMsg, err := anypb.New(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal message: %w", err)
	}
	anyMsg.TypeUrl = strings.TrimPrefix(anyMsg.TypeUrl, "type.googleapis.com/")
	anyMsg.TypeUrl = "monolith://" + anyMsg.TypeUrl
	return anyMsg, nil
}

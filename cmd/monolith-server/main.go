package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	monolith "reamde.dev/monolith/internal"
	"reamde.dev/monolith/proto/example"
)

type PingServer struct{}

func (s *PingServer) Ping(
	ctx context.Context,
	req *monolith.Request[example.PingRequest],
) (*monolith.Response[example.PingResponse], error) {
	res := &example.PingResponse{
		Nonce: req.Paylaod.Nonce,
	}
	return monolith.NewResponse(res), nil
}

type StringsServer struct{}

func (s *StringsServer) Upper(
	ctx context.Context,
	req *monolith.Request[example.UpperRequest],
) (*monolith.Response[example.UpperResponse], error) {
	res := &example.UpperResponse{
		Value: strings.ToUpper(req.Paylaod.Value),
	}
	return monolith.NewResponse(res), nil
}

func main() {
	keypair, err := monolith.GenerateKeyPair()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to generate key pair: %w", err))
	}

	transport := &monolith.TransportUTP{}
	listener, err := transport.Listen(
		context.Background(),
		"127.0.0.1:3000",
	)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to listen: %w", err))
	}

	sessions, err := monolith.NewSessionManager(
		transport,
		listener,
		keypair.PublicKey,
		keypair.PrivateKey,
	)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create session manager: %w", err))
	}

	pingServer := &PingServer{}
	err = example.RegisterPingService(sessions, pingServer)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to register ping service: %w", err))
	}

	stringsServer := &StringsServer{}
	err = example.RegisterStringsService(sessions, stringsServer)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to register strings service: %w", err))
	}

	select {}
}

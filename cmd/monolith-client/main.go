package main

import (
	"context"
	"fmt"
	"log"

	monolith "reamde.dev/monolith/internal"
	"reamde.dev/monolith/proto/example"
)

func main() {
	keypair, err := monolith.GenerateKeyPair()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to generate key pair: %w", err))
	}

	transport := &monolith.TransportUTP{}
	sessions, err := monolith.NewSessionManager(
		transport,
		nil,
		keypair.PublicKey,
		keypair.PrivateKey,
	)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create session manager: %w", err))
	}

	stringClient := example.NewStringsServiceClient(sessions)
	req := monolith.NewRequest(
		&example.UpperRequest{
			Value: "Jane",
		},
	)
	req.Target = &monolith.PeerAddr{
		Address:   "127.0.0.1:3000",
		Transport: "utp",
	}
	res, err := stringClient.Upper(context.Background(), req)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to call Upper: %w", err))
	}

	log.Println(res.Source.String())
	log.Println(res.Payload.Value)
}

// func main() {
// 	// construct monolith peer
// 	monolithConfig := monolith.NewConfig(
// 		monolith.WithBindAddress("tcp:localhost:8080"),
// 		monolith.WithPeerKey(os.Getenv("MONOLITH_PEER_KEY")),
// 		monolith.WithAccountKey(os.Getenv("MONOLITH_ACCOUNT_KEY")),
// 		monolith.WithBroker("tcp:broker.example.com:8080"),
// 		monolith.WithProvider("strings.*", "tcp:strings.example.com:8080"),
// 		monolith.WithProvider("ping.*", "tcp:ping.example.com:8080"),
// 	)

// 	// construct strings client
// 	stringsClient := exampleconnect.NewStringsServiceClient(monolithConfig)

// 	// do stuff
// 	// TODO: who does this request go to?
// 	res, err := stringsClient.Upper(
// 		context.Background(),
// 		monolith.NewRequest(
// 			&example.UpperRequest{
// 				Value: "Jane",
// 			},
// 		),
// 	)
// 	if err != nil {
// 		log.Fatal(fmt.Errorf("failed to call Upper: %w", err))
// 	}

// 	log.Println(res.Msg.Value)
// }

package main

import (
	"fmt"

	"reamde.dev/monolith/internal/broker"
	"reamde.dev/monolith/internal/client"
	"reamde.dev/monolith/internal/provider"
	"reamde.dev/monolith/proto/monolith"
)

func main() {
	brokerInfos := []*monolith.BrokerInfo{{
		Id: &monolith.Identity{
			Version: 1,
			Hash:    "broker1",
		},
		Peers: []*monolith.PeerInfo{{
			Addresses: []*monolith.Address{{
				Address: "127.0.0.1",
				Port:    8080,
			}},
		}},
	}}

	providerInfos := []*monolith.ProviderInfo{{
		Id: &monolith.Identity{
			Version: 1,
			Hash:    "provider1",
		},
		Topics: []string{"strings.ToUpper"},
		Peer: &monolith.PeerInfo{
			Addresses: []*monolith.Address{{
				Address: "127.0.0.1",
				Port:    8081,
			}},
		},
	}}

	accountInfo := &monolith.AccountInfo{
		Id: &monolith.Identity{
			Version: 1,
			Hash:    "client1",
		},
		Peer: &monolith.PeerInfo{
			Addresses: []*monolith.Address{{
				Address: "127.0.0.1",
				Port:    8082,
			}},
		},
		Brokers:   brokerInfos,
		Providers: providerInfos,
	}

	// Initialize and start broker
	brokerServer := &broker.BrokerServer{
		BrokerInfo:    brokerInfos[0],
		ProviderInfos: providerInfos,
	}
	brokerServer.Start("127.0.0.1:8080")

	// Initialize and start provider
	providerServer := &provider.ProviderServer{
		Info: providerInfos[0],
	}
	providerServer.Start("127.0.0.1:8081")

	// Initialize client and send request
	clientInstance := client.NewClient(accountInfo)
	resp, err := clientInstance.SendRequest("strings.ToUpper", []byte("hello"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if resp.Error != "" {
		fmt.Println("Error:", resp.Error)
		return
	}

	fmt.Println("Response:", string(resp.Body))
}

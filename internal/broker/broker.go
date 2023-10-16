package broker

import (
	"fmt"

	"reamde.dev/monolith/internal/rpc"
	"reamde.dev/monolith/proto/monolith"
)

// BrokerServer handles incoming requests from clients and forwards them to providers.
type BrokerServer struct {
	BrokerInfo    *monolith.BrokerInfo
	ProviderInfos map[string]*monolith.ProviderInfo
	rpcServer     *rpc.Server
}

func (b *BrokerServer) Handler(request rpc.Request) rpc.Response {
	provider, ok := b.ProviderInfos[request.Topic]
	if !ok {
		return rpc.Response{
			Error: fmt.Sprintf("No provider found for topic: %s", request.Topic),
		}
	}

	addr := provider.Peer.Addresses[0]
	host := fmt.Sprintf("%s:%d", addr.Address, addr.Port)

	client := &rpc.Client{}
	response, err := client.SendRequest(host, request)
	if err != nil {
		return rpc.Response{
			Error: fmt.Sprintf("Error forwarding to provider: %s", err),
		}
	}
	return response
}

func (b *BrokerServer) Start(bindAddress string) {
	b.rpcServer = &rpc.Server{
		Address: bindAddress,
		Handler: b.Handler,
	}
	go b.rpcServer.Listen() // Run in background
}

package provider

import (
	"strings"

	"reamde.dev/monolith/internal/rpc"
	"reamde.dev/monolith/proto/monolith"
)

// ProviderServer handles incoming requests and processes them.
type ProviderServer struct {
	Info      *monolith.ProviderInfo
	rpcServer *rpc.Server
}

func (p *ProviderServer) Handler(request rpc.Request) rpc.Response {
	// This is just a simple handler for demonstration. Depending on your requirements,
	// you can have more complex processing here.
	return rpc.Response{
		Body: []byte(strings.ToUpper(string(request.Body))),
	}
}

func (p *ProviderServer) Start(bindAddress string) {
	p.rpcServer = &rpc.Server{
		Address: bindAddress,
		Handler: p.Handler,
	}
	go p.rpcServer.Listen() // Run in background
}

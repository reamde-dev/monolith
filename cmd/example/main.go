package main

import (
	"context"
	"fmt"

	internal "reamde.dev/monolith/internal"
	"reamde.dev/monolith/internal/broker"
	"reamde.dev/monolith/internal/client"
	"reamde.dev/monolith/internal/provider"
	"reamde.dev/monolith/proto/example"
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

// RPC

// type (
// 	Request[T any] struct {
// 		Paylaod *T
// 	}
// 	Response[T any] struct {
// 		Payload *T
// 		Error   string
// 	}
// )

// type (
// 	RPCRequest struct {
// 		Path    string
// 		Payload []byte
// 	}
// 	RPCResponse struct {
// 		Payload []byte
// 		Error   string
// 	}
// )

// type HandlerFunc func(context.Context, *RPCRequest) (*RPCResponse, error)

// func NewHandler[Req, Res any](
// 	path string,
// 	handler func(context.Context, *Request[Req]) (*Response[Res], error),
// ) HandlerFunc {
// 	return func(ctx context.Context, rpcReq *RPCRequest) (*RPCResponse, error) {
// 		req := &Request[Req]{}
// 		err := json.Unmarshal(rpcReq.Payload, req)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to unmarshal request body: %w", err)
// 		}

// 		res, err := handler(ctx, req)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to handle request: %w", err)
// 		}

// 		resBytes, err := json.Marshal(res)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to marshal response body: %w", err)
// 		}

// 		return &RPCResponse{
// 			Payload: resBytes,
// 		}, nil
// 	}
// }

// type Handlers map[string]HandlerFunc

// type RPCServer struct {
// 	handlers Handlers
// }

// func (s *RPCServer) Handle(path string, handler HandlerFunc) {
// 	s.handlers[path] = handler
// }

// Generated server

// const (
// 	PingServicePingPath = "example.PingService/Ping"
// )

type PingService interface {
	Ping(context.Context, *internal.Request[example.PingRequest]) (*internal.Response[example.PingResponse], error)
}

// func NewPingServiceHandler(svc PingServiceHandler) (HandlerFunc, error) {
// 	pingServicePingHandler := NewHandler(PingServicePingPath, svc.Ping)
// 	return func(ctx context.Context, rpcReq *RPCRequest) (*RPCResponse, error) {
// 		switch rpcReq.Path {
// 		case PingServicePingPath:
// 			return pingServicePingHandler(ctx, rpcReq)
// 		default:
// 			return nil, fmt.Errorf("unknown procedure: %s", rpcReq.Path)
// 		}
// 	}, nil
// }

// func NewPingServiceHandlers(svc PingService) (internal.Handlers, error) {
// 	return internal.Handlers{
// 		"example.PingService": internal.NewHandler(svc.Ping),
// 	}, nil
// }

// func RegisterPingService(mgr *internal.SessionManager, svc PingService) error {
// 	handlers, err := NewPingServiceHandlers(svc)
// 	if err != nil {
// 		return fmt.Errorf("failed to create handlers: %w", err)
// 	}
// 	err = mgr.RegisterHandlers(handlers)
// 	if err != nil {
// 		return fmt.Errorf("failed to register handlers: %w", err)
// 	}
// 	return nil
// }

// Generated client

func NewPingServiceClient(mgr *internal.SessionManager) PingService {
	return &pingServiceClient{
		mgr: mgr,
	}
}

type pingServiceClient struct {
	mgr *internal.SessionManager
}

func (c *pingServiceClient) Ping(ctx context.Context, request *internal.Request[example.PingRequest]) (*internal.Response[example.PingResponse], error) {
	return internal.MakeRequest[example.PingRequest, example.PingResponse](c.mgr, ctx, "example.PingService/Ping", request)
}

// Implementation

type PingServiceImpl struct{}

func (s *PingServiceImpl) Ping(
	ctx context.Context,
	request *internal.Request[example.PingRequest],
) (*internal.Response[example.PingResponse], error) {
	return &internal.Response[example.PingResponse]{
		Payload: &example.PingResponse{
			Nonce: request.Paylaod.Nonce,
		},
	}, nil
}

// func someMain() {
// 	pingService := &PingServiceImpl{}
// 	pingHandler, err := NewPingServiceHandler(pingService)
// 	if err != nil {
// 		panic(err)
// 	}

// }

package client

import (
	"fmt"

	"reamde.dev/monolith/internal/rpc"
	"reamde.dev/monolith/proto/monolith"
)

type Client struct {
	Info      *monolith.ClientInfo
	rpcClient *rpc.Client
}

func NewClient(info *monolith.ClientInfo) *Client {
	return &Client{
		Info:      info,
		rpcClient: &rpc.Client{},
	}
}

func (c *Client) SendRequest(topic string, body []byte) (rpc.Response, error) {
	request := rpc.Request{
		Topic: topic,
		Body:  body,
	}
	addr := c.Info.Brokers[0].Peers[0].Addresses[0]
	host := fmt.Sprintf("%s:%d", addr.Address, addr.Port)
	return c.rpcClient.SendRequest(host, request)
}

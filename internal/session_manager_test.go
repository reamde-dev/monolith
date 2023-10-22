package monolith

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test for SessionManager
func TestSessionManager(t *testing.T) {
	srv, clt := newTestSessionManager(t)

	expRes := &RPCResponse{
		Payload: []byte("pong"),
	}
	handler := func(ctx context.Context, msg *RPCRequest) (*RPCResponse, error) {
		// err := msg.Respond(expRes.Document()) // TODO: reconsider .Respond
		// require.NoError(t, err)
		return expRes, nil
	}
	srv.RegisterHandler("test/ping", handler)

	req := &RPCRequest{
		Path:    "test/ping",
		Payload: []byte("ping"),
	}

	// dial the server
	// rpc, err := clt.dial(context.Background(), srv.PeerAddr())
	// require.NoError(t, err)

	// send a message
	// res, err := rpc.Request(context.Background(), req.Path, req.Payload)
	// require.NoError(t, err)
	// require.Equal(t, expRes, res)

	// make a request
	res, err := clt.Request(context.Background(), req, srv.PeerAddr())
	require.NoError(t, err)
	require.Equal(t, expRes, res)
}

func newTestSessionManager(t *testing.T) (srv *SessionManager, clt *SessionManager) {
	t.Helper()

	// create a new SessionManager for the server
	srvKeyPair, err := GenerateKeyPair()
	require.NoError(t, err)
	srvPub := srvKeyPair.PublicKey
	srvPrv := srvKeyPair.PrivateKey

	srvTransport := &TransportUTP{}
	srvListener, err := srvTransport.Listen(context.Background(), "127.0.0.1:0")
	require.NoError(t, err)
	srv, err = NewSessionManager(srvTransport, srvListener, srvPub, srvPrv)
	require.NoError(t, err)

	// create a new SessionManager for the client
	cltKeyPair, err := GenerateKeyPair()
	require.NoError(t, err)
	cltPub := cltKeyPair.PublicKey
	cltPrv := cltKeyPair.PrivateKey

	cltTransport := &TransportUTP{}
	cltListener, err := cltTransport.Listen(context.Background(), "127.0.0.1:0")
	require.NoError(t, err)
	clt, err = NewSessionManager(cltTransport, cltListener, cltPub, cltPrv)
	require.NoError(t, err)

	return srv, clt
}

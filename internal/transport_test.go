package monolith

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrapListener(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	require.NoError(t, err)
	defer ln.Close()

	pk := Must(GenerateKeyPair()).PublicKey

	// Wrap the dummy net.Listener in a listener
	wrapped := wrapListener(ln, "dummy", "some.host", pk)

	// Check that the wrapped listener has the correct PeerAddr
	expectedAddr := PeerAddr{
		Transport: "dummy",
		Address:   "some.host:8080",
		PublicKey: pk,
	}
	require.Equal(t, expectedAddr, wrapped.PeerAddr())

	// Check that the wrapped listener can be closed
	require.NoError(t, wrapped.Close())
}
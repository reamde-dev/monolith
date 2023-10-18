package monolith

import (
	"net"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSession_E2E_Pipe(t *testing.T) {
	messagePing := []byte("ping") // client to server
	messagePong := []byte("pong") // server to client

	// Create a server and a client that are connected to each other
	server, client := net.Pipe()

	// Generate the server's static keys
	serverKeyPair, err := GenerateKeyPair()
	serverPublicKey := serverKeyPair.PublicKey
	serverPrivateKey := serverKeyPair.PrivateKey
	require.NoError(t, err)

	// Generate the client's static keys
	clientKeyPair, err := GenerateKeyPair()
	clientPublicKey := clientKeyPair.PublicKey
	clientPrivateKey := clientKeyPair.PrivateKey
	require.NoError(t, err)

	// Perform the handshake from the server side
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		serverSession := NewSession(server)
		err = serverSession.DoServer(serverPublicKey, serverPrivateKey)
		require.NoError(t, err)

		// Receive the message from the server
		receivedMessage, err := serverSession.read()
		require.NoError(t, err)

		// Check that the received message is the same as the original message
		require.Equal(t, messagePing, receivedMessage)

		// Send a message from the server to the client
		_, err = serverSession.write(messagePong)
		require.NoError(t, err)

		// Done
		wg.Done()
	}()

	// Perform the handshake from the client side
	clientSession := NewSession(client)
	err = clientSession.DoClient(clientPublicKey, clientPrivateKey)
	require.NoError(t, err)

	// Send a message from the client to the server
	_, err = clientSession.write(messagePing)
	require.NoError(t, err)

	// Read the message from the server
	receivedMessage, err := clientSession.read()
	require.NoError(t, err)

	// Check that the received message is the same as the original message
	require.Equal(t, messagePong, receivedMessage)

	// Wait for the server to finish
	wg.Wait()
}

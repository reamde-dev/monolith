package monolith

import (
	"io"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type MockConn struct {
	Server net.Conn
	Client net.Conn
}

func NewMockConn() *MockConn {
	sr, cw := io.Pipe()
	cr, sw := io.Pipe()

	return &MockConn{
		Server: &MockConnEndpoint{
			Reader: sr,
			Writer: sw,
		},
		Client: &MockConnEndpoint{
			Reader: cr,
			Writer: cw,
		},
	}
}

type MockAddr struct {
	transport string
	address   string
}

func (m *MockAddr) Network() string {
	return m.transport
}

func (m *MockAddr) String() string {
	return m.address
}

type MockConnEndpoint struct {
	Reader io.Reader
	Writer io.Writer
}

func (m *MockConnEndpoint) Read(b []byte) (n int, err error) {
	return m.Reader.Read(b)
}

func (m *MockConnEndpoint) Write(b []byte) (n int, err error) {
	return m.Writer.Write(b)
}

func (m *MockConnEndpoint) Close() error {
	return nil
}

func (m *MockConnEndpoint) LocalAddr() net.Addr {
	return &MockAddr{
		transport: "mock",
		address:   "local",
	}
}

func (m *MockConnEndpoint) RemoteAddr() net.Addr {
	return &MockAddr{
		transport: "mock",
		address:   "remote",
	}
}

func (m *MockConnEndpoint) SetDeadline(t time.Time) error {
	return nil
}

func (m *MockConnEndpoint) SetReadDeadline(t time.Time) error {
	return nil
}

func (m *MockConnEndpoint) SetWriteDeadline(t time.Time) error {
	return nil
}

type MockSession struct {
	Server *Session
	Client *Session
}

func NewMockSession(t *testing.T) *MockSession {
	t.Helper()

	sr, cw := io.Pipe()
	cr, sw := io.Pipe()

	clientConn := &MockConnEndpoint{
		Reader: sr,
		Writer: sw,
	}

	serverConn := &MockConnEndpoint{
		Reader: cr,
		Writer: cw,
	}

	m := &MockSession{
		Server: NewSession(clientConn),
		Client: NewSession(serverConn),
	}

	serverKeyPair, err := GenerateKeyPair()
	serverPublic := serverKeyPair.PublicKey
	serverPrivate := serverKeyPair.PrivateKey
	require.NoError(t, err)

	clientKeyPair, err := GenerateKeyPair()
	clientPublic := clientKeyPair.PublicKey
	clientPrivate := clientKeyPair.PrivateKey
	require.NoError(t, err)

	serverDone := make(chan struct{})

	go func() {
		err := m.Server.DoServer(serverPublic, serverPrivate)
		require.NoError(t, err)
		close(serverDone)
	}()
	err = m.Client.DoClient(clientPublic, clientPrivate)
	require.NoError(t, err)

	<-serverDone
	return m
}

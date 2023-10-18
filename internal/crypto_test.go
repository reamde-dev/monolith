package monolith

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPublicKeyToBase58(t *testing.T) {
	kp0, err := GenerateKeyPair()
	require.NoError(t, err)

	b58 := kp0.PublicKey.String()

	pk1, err := ParsePublicKey(b58)
	require.NoError(t, err)

	require.True(t, kp0.PublicKey.Equal(pk1))
}

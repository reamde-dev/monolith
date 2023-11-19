package keystore

import (
	"crypto/sha256"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"reamde.dev/monolith/proto/monolith"
)

// Helper functions for testing
func setupFileStore(t *testing.T) *FileStore {
	tempDir, err := os.MkdirTemp("", "filestore_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}

	mockNewFunc := func() (*monolith.KeyPair, error) {
		return &monolith.KeyPair{
			PrivateKey: &monolith.PrivateKey{
				KeyType: &monolith.PrivateKey_Ed25519{
					Ed25519: &monolith.Ed25519Key{
						Key: []byte("private"),
					},
				},
			},
			PublicKey: &monolith.PublicKey{
				KeyType: &monolith.PublicKey_Ed25519{
					Ed25519: &monolith.Ed25519Key{
						Key: []byte("public"),
					},
				},
			},
		}, nil
	}

	mockHashFunc := func(kp protoreflect.ProtoMessage) (*monolith.Hash, error) {
		b, err := protojson.Marshal(kp)
		if err != nil {
			panic(fmt.Errorf("failed to marshal keypair: %w", err))
		}

		hash := sha256.New()
		_, err = hash.Write(b)
		if err != nil {
			panic(fmt.Errorf("failed to hash keypair: %w", err))
		}

		return &monolith.Hash{
			Type: monolith.Hash_TYPE_UNSPECIFIED,
			Hash: hash.Sum(nil),
		}, nil
	}

	fs := &FileStore{
		path: tempDir,
		new:  mockNewFunc,
		hash: mockHashFunc,
	}
	return fs
}

func TestPutAndGet(t *testing.T) {
	fs := setupFileStore(t)

	kp, _ := fs.new()
	hash, _ := fs.hash(kp)

	err := fs.Put(kp)
	require.NoError(t, err)

	got, err := fs.Get(hash)
	require.NoError(t, err)

	require.Equal(t, kp, got)
}

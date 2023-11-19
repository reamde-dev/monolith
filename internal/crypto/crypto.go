package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"

	"github.com/stackb/protoreflecthash"
	"google.golang.org/protobuf/reflect/protoreflect"

	"reamde.dev/monolith/proto/monolith"
)

func NewEd25519KeyPair() (*monolith.KeyPair, error) {
	sk, pk, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate key pair: %w", err)
	}

	return &monolith.KeyPair{
		PrivateKey: &monolith.PrivateKey{
			KeyType: &monolith.PrivateKey_Ed25519{
				Ed25519: &monolith.Ed25519Key{
					Key: sk,
				},
			},
		},
		PublicKey: &monolith.PublicKey{
			KeyType: &monolith.PublicKey_Ed25519{
				Ed25519: &monolith.Ed25519Key{
					Key: pk,
				},
			},
		},
	}, nil
}

var hasher = protoreflecthash.NewHasher(
	protoreflecthash.MessageFullnameIdentifier(),
	protoreflecthash.FieldNamesAsKeys(),
	protoreflecthash.SkipFields((monolith.Hash{}).Type.String()),
)

func Hash(msg protoreflect.ProtoMessage) (*monolith.Hash, error) {
	hash, err := hasher.HashProto(msg.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return &monolith.Hash{
		Type: monolith.Hash_TYPE_OBJECTHASH_SHA256,
		Hash: hash,
	}, nil
}

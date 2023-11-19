package main

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/stackb/protoreflecthash"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"

	"reamde.dev/monolith/proto/example"
	"reamde.dev/monolith/proto/monolith"
)

func main() {
	h1 := must(NewHash(must(NewAny(&example.TypeA{Data: "Hello, World!"}))))
	fmt.Printf("h1 hash: %x\n", h1.Hash)
	fmt.Printf("h1 type: %s\n", h1.Type)

	//

	h2 := must(NewHash(&example.TypeA{Data: "Hello, World!"}))
	fmt.Printf("h2 hash: %x\n", h2.Hash)
	fmt.Printf("h2 type: %s\n", h2.Type)

	//

	h3 := must(NewHash(&example.TypeB{Data: "Hello, World!"}))
	fmt.Printf("h3 hash: %x\n", h3.Hash)
	fmt.Printf("h3 type: %s\n", h3.Type)

	//

	msg := &monolith.Message{
		Payload: must(NewAny(&example.TypeA{Data: "Hello, World!"})),
	}

	hash := must(NewHash(msg))
	fmt.Printf("msg hash: %x\n", hash.Hash)
	fmt.Printf("msg hash type: %s\n", hash.Type)

	//

	kpc := must(NewEd25519KeyPair())
	fmt.Printf("public key (current): %x\n", kpc.PublicKey.GetEd25519().Key)

	kpn := must(NewEd25519KeyPair())
	fmt.Printf("public key (next): %x\n", kpn.PublicKey.GetEd25519().Key)

	sig := must(NewSignature(kpc, hash))
	fmt.Printf("signature: %x\n", sig.Signature)

	ok := must(VerifySignature(kpc.PublicKey, hash, sig))
	fmt.Printf("signature valid: %v\n", ok)
}

var hasher = protoreflecthash.NewHasher(
	protoreflecthash.MessageFullnameIdentifier(),
	protoreflecthash.FieldNamesAsKeys(),
	protoreflecthash.SkipFields((monolith.Hash{}).Algorithm.String()),
)

// Any is a wrapper around anypb.New that replaces the default type URL with
// a monolith specific one.
// https://developers.google.com/protocol-buffers/docs/proto3#any
func NewAny(msg protoreflect.ProtoMessage) (*anypb.Any, error) {
	anyMsg, err := anypb.New(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal message: %w", err)
	}
	anyMsg.TypeUrl = strings.TrimPrefix(anyMsg.TypeUrl, "type.googleapis.com/")
	anyMsg.TypeUrl = "monolith://" + anyMsg.TypeUrl
	return anyMsg, nil
}

func NewHash(msg protoreflect.ProtoMessage) (*monolith.Hash, error) {
	hash, err := hasher.HashProto(msg.ProtoReflect())
	if err != nil {
		return nil, err
	}
	msgType := string(msg.ProtoReflect().Descriptor().FullName())
	if msgAny, ok := msg.(*anypb.Any); ok {
		// any message types get prefixed with "type.googleapis.com/"
		// https://developers.google.com/protocol-buffers/docs/proto3#any
		msgType = strings.TrimPrefix(msgAny.TypeUrl, "type.googleapis.com/")
		// TODO(geoah): Should we be adding a monolith:// prefix here?
	}
	return &monolith.Hash{
		Algorithm: monolith.Hash_TYPE_OBJECTHASH_SHA256,
		Type:      msgType,
		Hash:      hash,
	}, nil
}

func NewKeygraph(
	currentKeyPair *monolith.KeyPair,
	nextKeyPair *monolith.KeyPair,
) (*monolith.Keygraph, error) {
	nh, err := NewHash(nextKeyPair)
	kg := &monolith.Keygraph{
		Version:  0,
		Sequence: 0,
		Event: &monolith.Keygraph_Rotation_{
			Rotation: &monolith.Keygraph_Rotation{
				Id:            uuid.New().String(),
				PublicKey:     currentKeyPair.PublicKey,
				NextPublicKey: nh,
			},
		},
	}
	return kg, err
}

func NewEd25519KeyPair() (*monolith.KeyPair, error) {
	pk, sk, err := ed25519.GenerateKey(rand.Reader)
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

func PrivateKeyFromProto(sk *monolith.PrivateKey) (ed25519.PrivateKey, error) {
	switch k := sk.KeyType.(type) {
	case *monolith.PrivateKey_Ed25519:
		return k.Ed25519.Key, nil
	default:
		return nil, fmt.Errorf("unsupported key type: %T", k)
	}
}

func PublicKeyFromProto(pk *monolith.PublicKey) (ed25519.PublicKey, error) {
	switch k := pk.KeyType.(type) {
	case *monolith.PublicKey_Ed25519:
		return k.Ed25519.Key, nil
	default:
		return nil, fmt.Errorf("unsupported key type: %T", k)
	}
}

func NewSignature(
	kp *monolith.KeyPair,
	hash *monolith.Hash,
) (*monolith.Signature, error) {
	sk, err := PrivateKeyFromProto(kp.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get private key: %w", err)
	}

	fp, err := NewHash(kp.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get fingerprint: %w", err)
	}

	return &monolith.Signature{
		Header: &monolith.Signature_Header{
			Algorthm:    monolith.Signature_TYPE_EDDSA,
			Fingerprint: fp,
		},
		Signature: ed25519.Sign(sk, hash.Hash),
	}, nil
}

func VerifySignature(
	ppk *monolith.PublicKey,
	hash *monolith.Hash,
	sig *monolith.Signature,
) (bool, error) {
	pk, err := PublicKeyFromProto(ppk)
	if err != nil {
		return false, fmt.Errorf("failed to get public key: %w", err)
	}

	fp, err := NewHash(ppk)
	if err != nil {
		return false, fmt.Errorf("failed to get fingerprint: %w", err)
	}

	if !bytes.Equal(fp.Hash, sig.Header.Fingerprint.Hash) {
		return false, fmt.Errorf("fingerprint mismatch")
	}

	return ed25519.Verify(pk, hash.Hash, sig.Signature), nil
}

func must[T any](t T, err error) T {
	if err != nil {
		panic(err.Error())
	}
	return t
}

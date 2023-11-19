package keystore

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"

	"reamde.dev/monolith/internal/crypto"
	"reamde.dev/monolith/proto/monolith"
)

type Store interface {
	// Get returns the keypair for the given key
	Get(fingerprint *monolith.Hash) (*monolith.KeyPair, error)
	// Put stores the given keypair
	Put(kp *monolith.KeyPair) error
	// List returns all keypairs in the store
	List() ([]*monolith.KeyPair, error)
	// Delete removes the keypair for the given key
	Delete(fingerprint *monolith.Hash) error
	// New creates a new keypair, stores it, and returns it
	New() (*monolith.KeyPair, error)
}

// FileStore implements the Store interface using a file system for storage.
type FileStore struct {
	path string
	new  func() (*monolith.KeyPair, error)
	hash func(protoreflect.ProtoMessage) (*monolith.Hash, error)
}

func NewFileStore(path string) *FileStore {
	return &FileStore{
		path: path,
		new:  crypto.NewEd25519KeyPair,
		hash: crypto.Hash,
	}
}

func (fs *FileStore) filename(fingerprint *monolith.Hash) string {
	return fmt.Sprintf("%s/%x", fs.path, fingerprint.Hash)
}

func (fs *FileStore) Get(fingerprint *monolith.Hash) (*monolith.KeyPair, error) {
	filePath := fs.filename(fingerprint)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var kp monolith.KeyPair
	err = protojson.Unmarshal(data, &kp)
	if err != nil {
		return nil, err
	}

	return &kp, nil
}

func (fs *FileStore) Put(kp *monolith.KeyPair) error {
	hash, err := fs.hash(kp)
	if err != nil {
		return err
	}

	filePath := fs.filename(hash)
	data, err := protojson.Marshal(kp)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

func (fs *FileStore) List() ([]*monolith.KeyPair, error) {
	files, err := os.ReadDir(fs.path)
	if err != nil {
		return nil, err
	}

	var keyPairs []*monolith.KeyPair
	for _, f := range files {
		data, err := os.ReadFile(filepath.Join(fs.path, f.Name()))
		if err != nil {
			return nil, err
		}

		var kp monolith.KeyPair
		err = protojson.Unmarshal(data, &kp)
		if err != nil {
			continue // skip files that can't be unmarshalled
		}

		keyPairs = append(keyPairs, &kp)
	}

	return keyPairs, nil
}

func (fs *FileStore) Delete(fingerprint *monolith.Hash) error {
	filePath := fs.filename(fingerprint)
	return os.Remove(filePath)
}

func (fs *FileStore) New() (*monolith.KeyPair, error) {
	kp, err := fs.new()
	if err != nil {
		return nil, err
	}

	err = fs.Put(kp)
	if err != nil {
		return nil, err
	}

	return kp, nil
}

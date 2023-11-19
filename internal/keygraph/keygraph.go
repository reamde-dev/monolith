package keygraph

import "reamde.dev/monolith/proto/monolith"

func New(kp *monolith.KeyPair) (*monolith.Keygraph, error) {
	return &monolith.Keygraph{
		Keys: []*monolith.KeyPair{kp},
	}, nil
}

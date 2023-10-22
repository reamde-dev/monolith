package monolith

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Shorthand string

const ShorthandPeerAddress Shorthand = "nimona://peer:addr:"

func (t Shorthand) String() string {
	return string(t)
}

type PeerAddr struct {
	_         string    `nimona:"$type,type=core/node.address"`
	Address   string    `nimona:"address,omitempty"`
	Transport string    `nimona:"transport,omitempty"`
	PublicKey PublicKey `nimona:"publicKey,omitempty"`
}

func ParsePeerAddr(addr string) (*PeerAddr, error) {
	regex := regexp.MustCompile(
		ShorthandPeerAddress.String() +
			`(?:([\w\d]+@))?([\w\d]+):([\w\d\.]+):(\d+)`,
	)
	matches := regex.FindStringSubmatch(addr)
	if len(matches) != 5 {
		return nil, errors.New("invalid input string")
	}

	publicKey := matches[1]
	transport := matches[2]
	host := matches[3]
	port := matches[4]

	a := &PeerAddr{}
	if publicKey != "" {
		publicKey = strings.TrimSuffix(publicKey, "@")
		key, err := ParsePublicKey(publicKey)
		if err != nil {
			return nil, fmt.Errorf("invalid public key, %w", err)
		}
		a.PublicKey = key
	}

	a.Transport = transport
	a.Address = fmt.Sprintf("%s:%s", host, port)

	return a, nil
}

func (a PeerAddr) String() string {
	b := strings.Builder{}
	b.WriteString(ShorthandPeerAddress.String())
	if a.PublicKey != nil {
		b.WriteString(a.PublicKey.String())
		b.WriteString("@")
	}
	b.WriteString(a.Transport)
	b.WriteString(":")
	b.WriteString(a.Address)
	return b.String()
}

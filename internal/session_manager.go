package monolith

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/golang-lru/v2/simplelru"
)

type (
	Request[T any] struct {
		Target  *PeerAddr
		Paylaod *T
	}
	Response[T any] struct {
		Source  PublicKey
		Payload *T
		Error   string
	}
)

func NewRequest[Req any](payload *Req) *Request[Req] {
	return &Request[Req]{
		Paylaod: payload,
	}
}

func NewResponse[Res any](payload *Res) *Response[Res] {
	return &Response[Res]{
		Payload: payload,
	}
}

type (
	RPCRequest struct {
		Path    string
		Target  *PeerAddr
		Payload []byte
	}
	RPCResponse struct {
		Source  PublicKey
		Payload []byte
		Error   string
	}
)

type HandlerFunc func(context.Context, *RPCRequest) (*RPCResponse, error)

func NewHandler[Req, Res any](
	handler func(context.Context, *Request[Req]) (*Response[Res], error),
) HandlerFunc {
	return func(ctx context.Context, rpcReq *RPCRequest) (*RPCResponse, error) {
		req := &Request[Req]{}
		err := json.Unmarshal(rpcReq.Payload, req)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal request body: %w", err)
		}

		res, err := handler(ctx, req)
		if err != nil {
			return nil, fmt.Errorf("failed to handle request: %w", err)
		}

		resBytes, err := json.Marshal(res)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response body: %w", err)
		}

		return &RPCResponse{
			Source:  res.Source,
			Payload: resBytes,
		}, nil
	}
}

type Handlers map[string]HandlerFunc

// SessionManager manages the dialing and accepting of connections.
// It maintains a cache of the last 100 connections.
type SessionManager struct {
	connCache  *simplelru.LRU[connCacheKey, *RPC]
	dialer     Dialer
	listener   Listener
	handlers   Handlers
	publicKey  PublicKey
	privateKey PrivateKey

	// TODO: check if we need these
	// resolver Resolver
	// aliases    xsync.Map[IdentityAlias, *IdentityInfo]
	// providers  xsync.Map[IdentityAlias, *IdentityInfo]
	// identities xsync.Map[KeygraphID, *IdentityInfo]
}

type connCacheKey struct {
	publicKeyInHex string
}

func NewSessionManager(
	dialer Dialer,
	listener Listener,
	publicKey PublicKey,
	privateKey PrivateKey,
) (*SessionManager, error) {
	connCache, err := simplelru.NewLRU(100, func(_ connCacheKey, ses *RPC) {
		err := ses.Close()
		if err != nil {
			// TODO: log error
			fmt.Println("error closing connection on eviction:", err)
			return
		}
	})
	if err != nil {
		return nil, fmt.Errorf("error creating connection cache: %w", err)
	}

	c := &SessionManager{
		connCache:  connCache,
		dialer:     dialer,
		listener:   listener,
		publicKey:  publicKey,
		privateKey: privateKey,
		handlers:   map[string]HandlerFunc{},
		// resolver:   &ResolverHTTP{},
	}

	if listener != nil {
		go func() {
			// nolint:errcheck // TODO: handle error
			c.handleConnections()
		}()
	}

	return c, nil
}

// dial dials the given address and returns a connection if successful. If the
// address is already in the cache, the cached connection is returned.
func (cm *SessionManager) dial(
	ctx context.Context,
	addr PeerAddr,
) (*RPC, error) {
	// check the cache
	existingConn, ok := cm.connCache.Get(cm.connCacheKey(addr.PublicKey))
	if ok {
		return existingConn, nil
	}

	// dial the address if it is not in the cache.
	conn, err := cm.dialer.Dial(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("error dialing %s: %w", addr, err)
	}

	// wrap the connection in a chunked connection
	ses := NewSession(conn)
	err = ses.DoServer(cm.publicKey, cm.privateKey)
	if err != nil {
		return nil, fmt.Errorf("error performing handshake: %w", err)
	}

	// if we have been given a public key for the remote peer, check it
	if addr.PublicKey != nil {
		if !ses.PublicKey().Equal(addr.PublicKey) {
			return nil, fmt.Errorf("public key mismatch")
		}
	}

	// wrap the session in an rpc connection
	rpc := NewRPC(ses)

	// start handling messages
	go func() {
		cm.handleSession(rpc)
	}()

	// add connection to cache
	cm.connCache.Add(cm.connCacheKey(ses.PublicKey()), rpc)

	return rpc, nil
}

func (cm *SessionManager) connCacheKey(k PublicKey) connCacheKey {
	return connCacheKey{
		publicKeyInHex: fmt.Sprintf("%x", k),
	}
}

// type RequestRecipientFn func(*requestRecipient)

// type requestRecipient struct {
// 	Alias      *IdentityAlias
// 	KeygraphID KeygraphID
// 	PeerAddr   *PeerAddr
// }

// func FromAlias(alias IdentityAlias) RequestRecipientFn {
// 	return func(r *requestRecipient) {
// 		r.Alias = &alias
// 	}
// }

// func FromIdentity(id KeygraphID) RequestRecipientFn {
// 	return func(r *requestRecipient) {
// 		r.KeygraphID = id
// 	}
// }

// func FromPeerAddr(peerAddr PeerAddr) RequestRecipientFn {
// 	return func(r *requestRecipient) {
// 		r.PeerAddr = &peerAddr
// 	}
// }

func (cm *SessionManager) Request(
	ctx context.Context,
	req *RPCRequest,
	rec *PeerAddr,
	// rfn RequestRecipientFn,
) (*RPCResponse, error) {
	return cm.requestFromPeerAddr(ctx, *rec, req)

	// rec := &requestRecipient{}
	// rfn(rec)

	// switch {
	// case rec.Alias != nil:
	// 	return cm.requestFromAlias(ctx, *rec.Alias, req)
	// case !rec.KeygraphID.IsEmpty():
	// 	return cm.requestFromIdentity(ctx, rec.KeygraphID, req)
	// case rec.PeerAddr != nil:
	// 	return cm.requestFromPeerAddr(ctx, *rec.PeerAddr, req)
	// default:
	// 	return nil, fmt.Errorf("no recipient specified")
	// }
}

// func (cm *SessionManager) requestFromAlias(
// 	ctx context.Context,
// 	alias IdentityAlias,
// 	req *Document,
// ) (*Response, error) {
// 	// resolve the alias
// 	info, err := cm.LookupAlias(alias)
// 	if err != nil {
// 		return nil, fmt.Errorf("error looking up alias %s: %w", alias, err)
// 	}

// 	return cm.requestFromIdentity(ctx, info.KeygraphID, req)
// }

// func (cm *SessionManager) requestFromIdentity(
// 	ctx context.Context,
// 	id KeygraphID,
// 	req *Document,
// ) (*Response, error) {
// 	// resolve the identity
// 	info, err := cm.LookupIdentity(id)
// 	if err != nil {
// 		return nil, fmt.Errorf("error looking up identity %s: %w", id, err)
// 	}

// 	return cm.requestFromPeerAddr(ctx, info.PeerAddresses[0], req)
// }

func (cm *SessionManager) requestFromPeerAddr(
	ctx context.Context,
	addr PeerAddr,
	req *RPCRequest,
) (*RPCResponse, error) {
	ses, err := cm.dial(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("error dialing %s: %w", addr, err)
	}

	res, err := ses.Request(ctx, req.Path, req.Payload)
	if err != nil {
		return nil, fmt.Errorf("error making request to %s: %w", addr, err)
	}

	return &RPCResponse{
		Source:  ses.session.remotePublicKey,
		Payload: res,
	}, nil
}

func (cm *SessionManager) handleConnections() error {
	errCh := make(chan error)
	// accept inbound connections.
	// if a connection with the same remote address already exists in the cache,
	// it is closed and removed before the new connection is added.
	go func() {
		for {
			conn, err := cm.listener.Accept()
			if err != nil {
				errCh <- fmt.Errorf("error accepting connection: %w", err)
				return
			}

			// start a new session, and perform the server side of the handshake
			// this will also perform the key exchange so after this we should
			// know the public key of the remote peer
			ses := NewSession(conn)
			err = ses.DoServer(cm.publicKey, cm.privateKey)
			if err != nil {
				// TODO: log error
				continue
			}

			// check if a connection with the same remote address already exists
			// in the cache.
			connCacheKey := cm.connCacheKey(ses.PublicKey())
			_, connectionExists := cm.connCache.Get(connCacheKey)
			if connectionExists {
				// remove the existing connection from the cache; this will
				// trigger the eviction callback which will close the connection
				cm.connCache.Remove(connCacheKey)
			}

			// wrap the session in an rpc connection
			rpc := NewRPC(ses)

			// start handling messages
			go func() {
				cm.handleSession(rpc)
			}()

			// add connection to cache
			cm.connCache.Add(connCacheKey, rpc)
		}
	}()

	return <-errCh
}

func (cm *SessionManager) handleSession(ses *RPC) {
	for {
		path, req, cb, err := ses.Read()
		if err != nil {
			// TODO log error
			fmt.Println("error reading message:", err)
			ses.Close() // TODO handle error
			return
		}

		// get the handler for the message type
		handler, ok := cm.handlers[path]
		if !ok {
			// TODO log error
			fmt.Println("no handler for path:", path)
			continue
		}

		// handle the message
		res, err := handler(context.Background(), &RPCRequest{
			Path:    path,
			Payload: req,
		})
		if err != nil {
			// TODO log error
			fmt.Println("error handling message:", err)
			continue
		}

		// send the response
		err = cb(res.Payload) // TODO: callback should be able to return an error
		if err != nil {
			// TODO log error
			fmt.Println("error sending response:", err)
			continue
		}
	}
}

func (cm *SessionManager) RegisterHandler(path string, handler HandlerFunc) error {
	fmt.Println("registering handler for path:", path)
	cm.handlers[path] = handler
	return nil
}

func (cm *SessionManager) RegisterHandlers(handlers Handlers) error {
	for path, handler := range handlers {
		cm.RegisterHandler(path, handler)
	}
	return nil
}

func (cm *SessionManager) PeerAddr() *PeerAddr {
	return &PeerAddr{
		Transport: cm.listener.PeerAddr().Transport,
		Address:   cm.listener.PeerAddr().Address,
		PublicKey: cm.publicKey,
	}
}

// Close closes all connections in the connection cache.
func (cm *SessionManager) Close() error {
	// purge will close all connections in the cache
	cm.connCache.Purge()
	return nil
}

// func (cm *SessionManager) LookupAlias(alias IdentityAlias) (*IdentityInfo, error) {
// 	if info, ok := cm.aliases.Load(alias); ok {
// 		return info, nil
// 	}

// 	identityInfo, err := cm.resolver.ResolveIdentityAlias(alias)
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to resolve provider alias: %w", err)
// 	}

// 	cm.aliases.Store(alias, identityInfo)
// 	cm.identities.Store(identityInfo.KeygraphID, identityInfo)

// 	// TODO add recursive lookup for user identities
// 	// TODO(geoah): fix "use"
// 	// if identityInfo.KeygraphID.Use == "provider" {
// 	// 	cm.providers.Store(alias, identityInfo)
// 	// }

// 	return identityInfo, nil
// }

// func (cm *SessionManager) LookupIdentity(id KeygraphID) (*IdentityInfo, error) {
// 	if info, ok := cm.identities.Load(id); ok {
// 		return info, nil
// 	}

// 	var identityInfo *IdentityInfo
// 	cm.providers.Range(func(key IdentityAlias, providerInfo *IdentityInfo) bool {
// 		ctx, cf := context.WithTimeout(context.Background(), time.Second)
// 		defer cf()

// 		for _, addr := range providerInfo.PeerAddresses {
// 			rctx := &RequestContext{}
// 			doc, err := RequestDocument(ctx, rctx, cm, id.DocumentID(), FromPeerAddr(addr))
// 			if err != nil {
// 				continue
// 			}

// 			err = identityInfo.FromDocument(doc)
// 			if err != nil {
// 				continue
// 			}

// 			return false
// 		}

// 		return true
// 	})

// 	if identityInfo == nil {
// 		return nil, fmt.Errorf("unable to resolve identity %s", id)
// 	}

// 	cm.identities.Store(id, identityInfo)

// 	// TODO verify the alias is indeed correct before storing or returning it
// 	if identityInfo.Alias.Hostname != "" {
// 		cm.aliases.Store(identityInfo.Alias, identityInfo)
// 	}

// 	// TODO(geoah): fix "use"
// 	// if identityInfo.KeygraphID.Use == "provider" {
// 	// 	cm.providers.Store(identityInfo.Alias, identityInfo)
// 	// }

// 	return identityInfo, nil
// }

func MakeRequest[Req, Res any](
	cm *SessionManager,
	ctx context.Context,
	path string,
	req *Request[Req],
) (*Response[Res], error) {
	rpcPayload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal rpc payload: %w", err)
	}

	rpcReq := &RPCRequest{
		Path:    path,
		Target:  req.Target,
		Payload: rpcPayload,
	}
	res, err := cm.Request(ctx, rpcReq, req.Target)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	var rpcRes Response[Res]
	err = json.Unmarshal(res.Payload, &rpcRes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	rpcRes.Source = res.Source

	return &rpcRes, nil
}

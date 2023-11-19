module reamde.dev/monolith

go 1.21.1

require (
	connectrpc.com/connect v1.11.1
	github.com/google/uuid v1.4.0
	github.com/hashicorp/golang-lru/v2 v2.0.7
	github.com/mr-tron/base58 v1.2.0
	github.com/neilalexander/utp v0.1.0
	github.com/oasisprotocol/curve25519-voi v0.0.0-20230904125328-1f23a7beb09a
	github.com/stackb/protoreflecthash v0.0.0-20230622204848-b7269c7fa663
	github.com/stretchr/testify v1.8.4
	golang.org/x/crypto v0.11.0
	golang.org/x/net v0.10.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/anacrolix/envpprof v1.3.0 // indirect
	github.com/anacrolix/missinggo v1.3.0 // indirect
	github.com/anacrolix/missinggo/perf v1.0.0 // indirect
	github.com/anacrolix/missinggo/v2 v2.5.1 // indirect
	github.com/anacrolix/sync v0.5.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/huandu/xstrings v1.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/stackb/protoreflecthash => ../go-protoreflecthash

// replace github.com/stackb/protoreflecthash => github.com/reamde-dev/go-protoreflecthash v0.0.0-20231115005445-54ecd5e6f900

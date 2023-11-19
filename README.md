# Monolith

An attempt at creating a "next generation" internet protocol, enabling
easy service creation and consumption across applications, fostering a
more connected and data-driven online experience.

## T

* [ ] Identities
* [ ] Routing
* [ ] Underlying data structures

## W

* Accounting
* Contacts
* Blogging
* 

## Development

```sh
go install github.com/bufbuild/buf/cmd/buf@latest
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

mkdir ./bin
go build -o bin/protoc-gen-monolith-go ./cmd/protoc-gen-monolith-go 
```

### Folder/Package Structure

The recommended folder/package structure for providers is:

```txt
/cmd
    /<provider>
        main.go
/rpc
    /<service>
        service.proto
        // auto-generated files
/internal
    /<service>server
        server.go
        // implementation
```

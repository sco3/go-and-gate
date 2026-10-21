# AND Gate grpc server

A lightweight Go gRPC service template demonstrating protocol buffer generation,
embedding of un-implemented stubs, and clean project layout.

## Project Structure

```text
.
├── cmd/
│   ├── client/     # gRPC client entrypoint
│   └── server/     # gRPC server entrypoint
├── pb/             # Generated Go stubs (protobuf & gRPC)
├── proto/          # Source .proto schema definitions
├── Justfile        # Task runner for code generation
└── go.mod          # Go module definition
```

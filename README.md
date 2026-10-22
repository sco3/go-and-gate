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

TBD:

```text
Dockerfile:1:1
  1:1  high  Ensure that HEALTHCHECK instructions have been added to container images  checkov/CKV_DOCKER_2
  1:1  high  Ensure that a user for the container has been created                     checkov/CKV_DOCKER_3
 12:1  high  Ensure the base image uses a non latest version tag                       checkov/CKV_DOCKER_7
```

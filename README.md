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

terraform/main.tf:15:1
 15:1  high  Image Tag should be fixed - not latest or blank                   checkov/CKV_K8S_14
 15:1  high  Image Pull Policy should be Always                                checkov/CKV_K8S_15
 15:1  high  The default namespace should not be used                          checkov/CKV_K8S_21
 15:1  high  Minimize the admission of containers with the NET_RAW capability  checkov/CKV_K8S_28
 15:1  high  Apply security context to your pods, deployments and daemon_sets  checkov/CKV_K8S_29
 15:1  high  Apply security context to your pods and containers                checkov/CKV_K8S_30
 15:1  high  Image should use digest                                           checkov/CKV_K8S_43
 15:1  high  Liveness Probe Should be Configured                               checkov/CKV_K8S_8
 15:1  high  Readiness Probe Should be Configured                              checkov/CKV_K8S_9
 66:1  high  The default namespace should not be used                          checkov/CKV_K8S_21

```

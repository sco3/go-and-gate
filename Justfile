# List targets
default:
    just --list

# Generate gRPC code from proto files
proto:
    mkdir -p pb
    protoc \
      --proto_path=proto \
      --go_out=pb --go_opt=paths=source_relative \
      --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
      proto/and-gate.proto

# Start server
serve:
    go build 
    ./and-gate -start
    
# Smoke test with curl 
ping:
    curl -v --http2-prior-knowledge \
      -X POST \
      -H "Content-Type: application/grpc" \
      http://localhost:5555/time.StartService/GetStartNanos

# Smoke test with client
pingc:
    go build 
    ./and-gate -get

docker:
    docker build -t and-gate-server:latest .
    docker images
    
docker-serve:
    docker rm -f and-gate
    docker run -d --name and-gate -p 5555:5555 and-gate-server:latest 
    
kind-create:
    kind create cluster --name and-gate-cluster
    
kind-delete:
    kind delete cluster --name and-gate-cluster

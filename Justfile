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
    ./and-gate -address :5555
    
# Smoke test with curl 
ping:
    curl -v --http2-prior-knowledge \
      -X POST \
      -H "Content-Type: application/grpc" \
      http://localhost:5555/time.StartService/GetStartNanos

# Smoke test with client

pingc:
    go build ./util/client
    ./util/client/client
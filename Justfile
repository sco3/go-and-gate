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

it:
   docker rm -f redis-server 2>/dev/null|| true
   docker run -d --name redis-server -p 6379:6379 redis
   go build
   EXPECTED_REPLICAS=2 ./and-gate -start -host one -address :5551 &
   EXPECTED_REPLICAS=2 ./and-gate -start -host two -address :5552 &
   sleep .3
   ./and-gate -get -address :5551
   ./and-gate -get -address :5552
   ./and-gate -stop -address :5551
   ./and-gate -stop -address :5552

   docker rm -f redis-server

docker-image:
    docker build -t and-gate-server:latest .
    docker images
    
docker-serve:
    docker rm -f and-gate
    docker run -d --name and-gate -p 5555:5555 and-gate-server:latest 
    
kind-create:
    kind create cluster --name and-gate-cluster
    
kind-delete:
    kind delete cluster --name and-gate-cluster

kind-load-image:
    kind load docker-image and-gate-server:latest --name and-gate-cluster
kube-port-forward:
    kubectl port-forward svc/and-gate-service 5555:5555
    
terr-init:
    cd terraform &&     terraform int

terr-plan:
    cd terraform &&     terraform plan
    
terr-apply:
    cd terraform &&     terraform apply --auto-approve
    
terr-destroy:
    cd terraform &&     terraform destroy --auto-approve
    
    
kube-pingc:
    kubectl delete pod client-test --ignore-not-found=true
    kubectl run client-test --image=and-gate-server:latest --image-pull-policy=IfNotPresent -- -start=false -get -address and-gate-service:5555
    sleep 2
    kubectl logs client-test
    kubectl delete pod client-test --ignore-not-found=true
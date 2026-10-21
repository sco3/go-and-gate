


go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ln -sf ~/go/bin/protoc-gen-go ~/bin/
ln -sf ~/go/bin/protoc-gen-go-grpc ~/bin/


rm -rf pb
mkdir -p pb
protoc \
  --go_out=pb --go_opt=paths=source_relative \
  --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
  and-gate.proto

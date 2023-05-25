gRPC is a "high-performance open source universal RPC framework."

Let's build a caching service together using gRPC.



## packages
 1. go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
 2.  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
 
# Generate proto file
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/app.proto 
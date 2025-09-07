go get google.golang.org/grpc
go get google.golang.org/protobuf

Generate Go code from proto

Run:

protoc --go_out=. --go-grpc_out=. proto/user.proto

This will generate user.pb.go and user_grpc.pb.go under proto/userpb.

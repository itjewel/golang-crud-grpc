go get google.golang.org/grpc
go get google.golang.org/protobuf

Generate Go code from proto

Run:

protoc --go_out=. --go-grpc_out=. proto/user.proto

This will generate user.pb.go and user_grpc.pb.go under proto/userpb.

প্রথমে folder create করো:

mkdir -p proto/userpb

-p দিয়ে parent folders না থাকলেও create হবে।

তারপর generate করো:

protoc --go_out=./proto/userpb --go-grpc_out=./proto/userpb proto/user.proto

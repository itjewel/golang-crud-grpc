package main

import (
	"log"
	"net"
	"net/http"

	"golang-crud/database"
	"golang-crud/grpc"
	"golang-crud/routes"

	pb "golang-crud/proto/userpb"

	"google.golang.org/grpc"
)

func main() {
    // DB connection
    database.Connect()
    defer database.DB.Close()

    // HTTP server
    mux := http.NewServeMux()
    routes.CategoryRoutes(mux)
    routes.ProductRoutes(mux)
    routes.UserRoutes(mux)

    go func() {
        log.Println("HTTP Server running at :8000")
        if err := http.ListenAndServe(":8000", mux); err != nil {
            log.Fatal(err)
        }
    }()

    // gRPC server
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer() // <-- use grpc.NewServer
    pb.RegisterUserServiceServer(grpcServer, &YourUserServer{}) // replace with your server struct

    log.Println("gRPC Server running at :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve gRPC: %v", err)
    }
}

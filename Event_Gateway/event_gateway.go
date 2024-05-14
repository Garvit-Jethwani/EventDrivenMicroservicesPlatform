
package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

type server struct {}

func main() {
    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    srv := grpc.NewServer()
    reflection.Register(srv)
    if err := srv.Serve(listener); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

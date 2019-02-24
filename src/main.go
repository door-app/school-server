package main

import (
    "google.golang.org/grpc"
    "log"
    "net"
)

func main() {
    listen, err := net.Listen("tcp", ":9090")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    if err := s.Serve(listen); err != nil {
        log.Fatalf("failded to serve: %v", err)
    }
}

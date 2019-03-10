package main

import (
    pb "door-app.net/school-server/protoc"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "log"
    "net"
)

type server struct{}

func (s *server) SayHello(context.Context, *pb.HelloRequest) (*pb.HelloReply, error) {
    // panic("implement me")
    return &pb.HelloReply{Message: "hello"}, nil
}

func main() {
    listen, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterGreeterServer(s, &server{})
    reflection.Register(s)
    if err := s.Serve(listen); err != nil {
        log.Fatalf("failded to serve: %v", err)
    }
}

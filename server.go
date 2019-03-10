package main

import (
  "fmt"
  "github.com/southernvg/go-protobuf-lambda-pb"
  "google.golang.org/grpc"
  "log"
  "net"
  "context"
)

// For gRPC
type server struct {}

func (s *server) Ping(ctx context.Context, in *ping_service.PingRequest) (*ping_service.PingResponse, error) {
  log.Println("Got request from client: " + in.Text)
  return &ping_service.PingResponse{Text: in.Text}, nil
}

func main() {
  var err error

  lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }

  grpcServer := grpc.NewServer()

  ping_service.RegisterPingServer(grpcServer, &server{})

  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %s", err)
  }

}

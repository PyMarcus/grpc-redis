package main

import (
	"context"
	"log"
	"net"

	"github.com/PyMarcus/gRPC-redis/internal/grpc"
	"github.com/PyMarcus/gRPC-redis/internal/repository"
	"github.com/PyMarcus/gRPC-redis/internal/utils"
	"github.com/PyMarcus/gRPC-redis/proto"
	g "google.golang.org/grpc"
)

func main() {
	utils.Init()

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Fail to start server: %v", err)
	}

	_context := context.Background()

	redisRepository := repository.NewRedisRepository("redis-grpc", "6379")
	redisRepository.Connect(_context)

	server := grpc.NewGRPCServer(redisRepository)

	s := g.NewServer()
	proto.RegisterKVStoreServer(s, server)
	
	log.Println("Starting server...")
	if err := s.Serve(listen); err != nil{
		log.Println("Fail to listen server...", err)
	}

}

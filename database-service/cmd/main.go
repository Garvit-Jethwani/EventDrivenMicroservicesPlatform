package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Garvit-Jethwani/database-service/config"
	"github.com/Garvit-Jethwani/database-service/database"
	"github.com/Garvit-Jethwani/database-service/proto"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create database service
	databaseService, err := database.NewDatabaseService(cfg)
	if err != nil {
		log.Fatalf("Failed to create database service: %v", err)
	}

	// Create gRPC server
	lis, err := net.Listen("tcp", cfg.GRPCAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterDatabaseServiceServer(s, databaseService)
	reflection.Register(s)

	log.Printf("Starting Database Service on %s", cfg.GRPCAddress)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

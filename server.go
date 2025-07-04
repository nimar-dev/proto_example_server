package main

import (
	"context"
	"log"
	"net"

	pb "github.com/nimar-dev/proto_example_server/coffeeshop_protos"
	"google.golang.org/grpc"
)

type myServer struct {
	pb.UnimplementedCoffeeShopNimarServer
}

// GetMenu implements CoffeeShopNimarServer
func (s *myServer) GetMenu(ctx context.Context, req *pb.MenuRequest) (*pb.MenuResponse, error) {
	// Here we would normally fetch the menu from a database or another service.
	// For this example, we will return a static menu.
	menu := &pb.MenuResponse{
		Items: []*pb.MenuItem{
			{ItemName: "Espresso", Price: 2.50},
			{ItemName: "Latte", Price: 3.00},
			{ItemName: "Cappuccino", Price: 3.50},
			{ItemName: "Americano", Price: 2.75},
			{ItemName: "Mocha", Price: 3.75},
		},
	}
	return menu, nil
}

// register the server with gRPC
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCoffeeShopNimarServer(grpcServer, &myServer{})
	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

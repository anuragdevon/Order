package server

import (
	"fmt"
	"log"
	"net"
	"order/pkg/client"
	"order/pkg/config"
	"order/pkg/pb"
	"order/pkg/repository"
	"order/pkg/services"

	"google.golang.org/grpc"
)

func Run() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := &repository.Database{}
	err = db.Connect(&c)
	if err != nil {
		log.Panic("failed to connect to database:", err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	inventorySvc := client.InitInventoryServiceClient(c.InventorySvcUrl)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Order Svc on", c.Port)

	newOrderService := services.NewOrderService(db.DB, &inventorySvc)
	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, newOrderService)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

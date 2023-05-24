package client

import (
	"context"
	"fmt"
	"order/pkg/pb"

	"google.golang.org/grpc"
)

func InitInventoryServiceClient(url string) InventoryServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := InventoryServiceClient{
		Client: pb.NewInventoryServiceClient(cc),
	}

	return c
}

type InventoryServiceClient struct {
	Client pb.InventoryServiceClient
}

func (c *InventoryServiceClient) GetItem(productId int64) (*pb.GetItemResponse, error) {
	req := &pb.GetItemRequest{
		Id: productId,
	}

	return c.Client.GetItem(context.Background(), req)
}

func (c *InventoryServiceClient) UpdateItem(id, quantity int64) (*pb.UpdateItemResponse, error) {
	req := &pb.UpdateItemRequest{
		Id:       id,
		Quantity: quantity,
	}

	return c.Client.UpdateItem(context.Background(), req)
}

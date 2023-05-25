package client

import (
	"context"
	"fmt"
	"order/pkg/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitInventoryServiceClient(url string) InventoryServiceClient {
	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

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

func (c *InventoryServiceClient) GetItem(itemId int64) (*pb.GetItemResponse, error) {
	req := &pb.GetItemRequest{
		Id: itemId,
	}

	return c.Client.GetItem(context.Background(), req)
}

func (c *InventoryServiceClient) DecreaseItemQuantity(id, quantity int64) (*pb.DecreaseItemQuantityResponse, error) {
	req := &pb.DecreaseItemQuantityRequest{
		Id:       id,
		Quantity: quantity,
	}

	return c.Client.DecreaseItemQuantity(context.Background(), req)
}

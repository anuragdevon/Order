package client

import (
	"context"
	"order/pkg/pb"
)

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

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

func (c *InventoryServiceClient) UpdateItem(id, quantity, price int64, name string) (*pb.UpdateItemResponse, error) {
	req := &pb.UpdateItemRequest{
		Id:       id,
		Quantity: quantity,
		Price:    price,
		Name:     name,
	}

	return c.Client.UpdateItem(context.Background(), req)
}

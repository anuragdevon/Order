package mocks

import (
	"context"

	"order/pkg/pb"

	"google.golang.org/grpc"
)

type MockInventoryServiceClient struct {
	pb.InventoryServiceClient
}

func (m *MockInventoryServiceClient) GetItem(ctx context.Context, req *pb.GetItemRequest, opts ...grpc.CallOption) (*pb.GetItemResponse, error) {
	response := &pb.GetItemResponse{
		Status: 200,
		Data:   &pb.GetItemData{Id: req.Id, Name: "Item Name", Quantity: 10, Price: 100},
	}

	return response, nil
}

func (m *MockInventoryServiceClient) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest, opts ...grpc.CallOption) (*pb.UpdateItemResponse, error) {
	response := &pb.UpdateItemResponse{
		Status: 200,
		Data:   &pb.GetItemData{Id: req.Id, Name: req.Name, Quantity: req.Quantity, Price: req.Price},
	}

	return response, nil
}

package services

import (
	"context"
	"net/http"

	"order/pkg/client"
	"order/pkg/pb"
	"order/pkg/repository"
	"order/pkg/repository/models"

	"gorm.io/gorm"
)

type OrderService struct {
	db           *gorm.DB
	InventorySvc client.InventoryServiceClient
	pb.UnimplementedOrderServiceServer
}

func NewOrderService(db *gorm.DB, inventorysvc *client.InventoryServiceClient) *OrderService {
	return &OrderService{
		db:                              db,
		InventorySvc:                    *inventorysvc,
		UnimplementedOrderServiceServer: pb.UnimplementedOrderServiceServer{},
	}
}

func (os *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	itemResp, err := os.InventorySvc.GetItem(req.ItemId)

	if err != nil {
		return &pb.CreateOrderResponse{
				Status: http.StatusBadGateway,
				Error:  err.Error(),
			},
			nil
	} else if itemResp.Status >= http.StatusNotFound {
		return &pb.CreateOrderResponse{
				Status: itemResp.Status,
				Error:  itemResp.Error,
			},
			nil
	} else if itemResp.Data.Quantity < req.Quantity {
		return &pb.CreateOrderResponse{
				Status: http.StatusConflict,
				Error:  "Quantity is insufficient",
			},
			nil
	}

	order := models.Order{
		ItemId: itemResp.Data.Id,
		UserId: req.UserId,
	}
	db := repository.Database{DB: os.db}

	err = db.CreateOrder(&order)

	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: err.Error()}, nil
	}
	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.Id,
	}, nil
}

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
	db := repository.Database{DB: os.db}

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
		ItemId:   itemResp.Data.Id,
		UserId:   req.UserId,
		Quantity: req.Quantity,
	}

	err = db.CreateOrder(&order)

	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: err.Error()}, nil
	}

	_, err = os.InventorySvc.DecreaseItemQuantity(order.ItemId, order.Quantity)

	if err != nil {
		db.DeleteOrder(order.Id)
		return &pb.CreateOrderResponse{Status: http.StatusInternalServerError, Error: "Internal server error"}, nil

	}
	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.Id,
	}, nil
}

func (os *OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	db := repository.Database{DB: os.db}

	order, err := db.GetOrder(req.Id, req.UserId)

	if err != nil {
		return &pb.GetOrderResponse{Status: http.StatusInternalServerError, Error: err.Error()}, nil
	}

	if order == nil {
		return &pb.GetOrderResponse{Status: http.StatusNotFound, Error: "Order not found"}, nil
	}

	itemResp, err := os.InventorySvc.GetItem(order.ItemId)
	if err != nil {
		return &pb.GetOrderResponse{Status: http.StatusBadGateway, Error: err.Error()}, nil
	}

	getOrderData := &pb.GetOrderData{
		Id:       order.Id,
		ItemId:   itemResp.Data.Id,
		Name:     itemResp.Data.Name,
		Quantity: order.Quantity,
		Price:    itemResp.Data.Price,
	}

	return &pb.GetOrderResponse{
		Status: http.StatusOK,
		Data:   getOrderData,
	}, nil
}

func (os *OrderService) GetAllOrder(ctx context.Context, req *pb.GetAllOrdersRequest) (*pb.GetAllOrdersResponse, error) {
	db := repository.Database{DB: os.db}

	orders, _ := db.GetOrdersByUserID(req.UserId)

	var getOrderData []*pb.GetAllOrdersData
	for _, order := range orders {
		orderData := &pb.GetAllOrdersData{
			Id:       order.Id,
			ItemId:   order.ItemId,
			Quantity: order.Quantity,
		}

		getOrderData = append(getOrderData, orderData)
	}

	return &pb.GetAllOrdersResponse{
		Status: http.StatusOK,
		Data:   getOrderData,
	}, nil
}

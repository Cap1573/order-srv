package handler

import (
	"context"
	order "github.com/Cap1573/order-srv/proto/order"
)

type Order struct{}

func (h *Order) FindOrderByUserId(ctx context.Context, req *order.UserId, rsp *order.Orders) error {
	od:=&order.OrderDetail{OrderId:1,OrderNum:"1"}
	rsp.Orders = append(rsp.Orders, od)
	return nil
}



package main

import (
	"context"
	"fmt"
	go_micro_srv_order "github.com/Cap1573/order-srv/proto/order"
	"github.com/micro/go-micro"
)

func main() {
	testOrder:=go_micro_srv_order.NewOrderService("go.micro.srv.order",micro.NewService().Client())
	rsp,err:=testOrder.FindOrderByUserId(context.TODO(),&go_micro_srv_order.UserId{UserId:"1234"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp)
}

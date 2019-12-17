package main

import (
	"github.com/Cap1573/order-srv/handler"
	order "github.com/Cap1573/order-srv/proto/order"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.order"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler

	order.RegisterOrderHandler(service.Server(), new(handler.Order))


	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/diskordanz/consumer/http/server"

	cHandler "github.com/diskordanz/consumer/pkg/category/handler"
	uHandler "github.com/diskordanz/consumer/pkg/consumer/handler"
	fHandler "github.com/diskordanz/consumer/pkg/franchise/handler"
	lHandler "github.com/diskordanz/consumer/pkg/location/handler"
	oHandler "github.com/diskordanz/consumer/pkg/order/handler"
	pHandler "github.com/diskordanz/consumer/pkg/product/handler"

	"github.com/diskordanz/consumer/service"

	cStorage "github.com/diskordanz/consumer/pkg/category/storage"
	uStorage "github.com/diskordanz/consumer/pkg/consumer/storage"
	fStorage "github.com/diskordanz/consumer/pkg/franchise/storage"
	lStorage "github.com/diskordanz/consumer/pkg/location/storage"
	oStorage "github.com/diskordanz/consumer/pkg/order/storage"
	pStorage "github.com/diskordanz/consumer/pkg/product/storage"
)

func main() {
	franchiseStorage := fStorage.NewInMemoryFranchiseStorage()
	categoryStorage := cStorage.NewInMemoryCategoryStorage()
	locationStorage := lStorage.NewInMemoryLocationStorage()
	orderStorage := oStorage.NewInMemoryOrderStorage()
	productStorage := pStorage.NewInMemoryProductStorage()
	consumerStorage := uStorage.NewInMemoryConsumerStorage()

	franchiseHandler := fHandler.NewFranchiseHandler(franchiseStorage)
	categoryHandler := cHandler.NewCategoryHandler(categoryStorage)
	locationHandler := lHandler.NewLocationHandler(locationStorage)
	orderHandler := oHandler.NewOrderHandler(orderStorage)
	productHandler := pHandler.NewProductHandler(productStorage)
	consumerHandler := uHandler.NewConsumerHandler(consumerStorage)

	srv := service.NewConsumerService(franchiseHandler, locationHandler, categoryHandler, productHandler, orderHandler, consumerHandler)

	api := server.NewAPI(&srv)
	api.AssignRoutes()
	api.Run(":8080")
}

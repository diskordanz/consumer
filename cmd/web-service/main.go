package main

import (
	"github.com/diskordanz/consumer/http/server"

	cHandler "github.com/diskordanz/consumer/pkg/category/handler"
	fHandler "github.com/diskordanz/consumer/pkg/franchise/handler"
	lHandler "github.com/diskordanz/consumer/pkg/location/handler"
	oHandler "github.com/diskordanz/consumer/pkg/order/handler"
	pHandler "github.com/diskordanz/consumer/pkg/product/handler"
	uHandler "github.com/diskordanz/consumer/pkg/profile/handler"

	"github.com/diskordanz/consumer/service"

	cStorage "github.com/diskordanz/consumer/pkg/category/storage"
	fStorage "github.com/diskordanz/consumer/pkg/franchise/storage"
	lStorage "github.com/diskordanz/consumer/pkg/location/storage"
	oStorage "github.com/diskordanz/consumer/pkg/order/storage"
	pStorage "github.com/diskordanz/consumer/pkg/product/storage"
	uStorage "github.com/diskordanz/consumer/pkg/profile/storage"
)

func main() {
	franchiseStorage := fStorage.NewInMemoryFranchiseStorage()
	categoryStorage := cStorage.NewInMemoryCategoryStorage()
	locationStorage := lStorage.NewInMemoryLocationStorage()
	orderStorage := oStorage.NewInMemoryOrderStorage()
	productStorage := pStorage.NewInMemoryProductStorage()
	profileStorage := uStorage.NewInMemoryProfileStorage()

	franchiseHandler := *fHandler.NewFranchiseHandler(franchiseStorage)
	categoryHandler := *cHandler.NewCategoryHandler(categoryStorage)
	locationHandler := *lHandler.NewLocationHandler(locationStorage)
	orderHandler := *oHandler.NewOrderHandler(orderStorage)
	productHandler := *pHandler.NewProductHandler(productStorage)
	profileHandler := *uHandler.NewProfileHandler(profileStorage)

	srv := service.NewConsumerService(franchiseHandler, locationHandler, categoryHandler, productHandler, orderHandler, profileHandler)

	api := server.NewAPI(&srv)
	api.AssignRoutes()
	api.Run(":8080")
}

package main

import (
	"log"
	"os"

	"github.com/diskordanz/consumer/http/server"
	"google.golang.org/grpc"

	cHandler "github.com/diskordanz/consumer/pkg/category/handler"
	uHandler "github.com/diskordanz/consumer/pkg/consumer/handler"
	fHandler "github.com/diskordanz/consumer/pkg/franchise/handler"
	lHandler "github.com/diskordanz/consumer/pkg/location/handler"
	oHandler "github.com/diskordanz/consumer/pkg/order/handler"
	pHandler "github.com/diskordanz/consumer/pkg/product/handler"

	"github.com/diskordanz/consumer/service"
	"upper.io/db.v3/postgresql"

	cStorage "github.com/diskordanz/consumer/pkg/category/storage"
	uStorage "github.com/diskordanz/consumer/pkg/consumer/storage"
	fStorage "github.com/diskordanz/consumer/pkg/franchise/storage"
	lStorage "github.com/diskordanz/consumer/pkg/location/storage"
	oStorage "github.com/diskordanz/consumer/pkg/order/storage"
	pStorage "github.com/diskordanz/consumer/pkg/product/storage"
	pb "github.com/iamnotjustice/web-metrics/pkg/api"
)

var dbSettings = postgresql.ConnectionURL{
	Database: os.Getenv("DB_NAME"),
	Host:     os.Getenv("DB_HOST"),
	User:     os.Getenv("DB_USERNAME"),
	Password: os.Getenv("DB_PASSWORD"),
}

func main() {

	db, err := postgresql.Open(dbSettings)

	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()

	franchiseStorage := fStorage.NewUpperFranchiseStorage(db)
	categoryStorage := cStorage.NewUpperCategoryStorage(db)
	locationStorage := lStorage.NewUpperLocationStorage(db)
	orderStorage := oStorage.NewUpperOrderStorage(db)
	productStorage := pStorage.NewUpperProductStorage(db)
	consumerStorage := uStorage.NewUpperConsumerStorage(db)

	franchiseHandler := fHandler.NewFranchiseHandler(franchiseStorage)
	categoryHandler := cHandler.NewCategoryHandler(categoryStorage)
	locationHandler := lHandler.NewLocationHandler(locationStorage)
	orderHandler := oHandler.NewOrderHandler(orderStorage)
	productHandler := pHandler.NewProductHandler(productStorage)
	consumerHandler := uHandler.NewConsumerHandler(consumerStorage)

	srv := service.NewConsumerService(franchiseHandler, locationHandler, categoryHandler, productHandler, orderHandler, consumerHandler)
	conn, err := grpc.Dial(os.Getenv("METRICS_HOST"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error dialing to client: %+v", err)
	}
	defer conn.Close()
	metricsService := pb.NewMetricsServiceClient(conn)

	api := server.NewAPI(&srv, &metricsService)
	api.AssignRoutes()
	api.Run(os.Getenv("CONSUMER_HOST"))
}

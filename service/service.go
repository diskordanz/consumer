package service

import (
	cHandler "github.com/diskordanz/consumer/pkg/category/handler"
	fHandler "github.com/diskordanz/consumer/pkg/franchise/handler"
	lHandler "github.com/diskordanz/consumer/pkg/location/handler"
	oHandler "github.com/diskordanz/consumer/pkg/order/handler"
	pHandler "github.com/diskordanz/consumer/pkg/product/handler"
	uHandler "github.com/diskordanz/consumer/pkg/profile/handler"
	"github.com/diskordanz/consumer/service/model"
)

type Service interface {
	GetFranchise(int) (model.Franchise, error)
	ListFranchises(count, offset int) ([]model.Franchise, error)
	SearchFranchisesByName(count, offset int, name string) ([]model.Franchise, error)

	GetLocationsOfFranchise(franchiseID, count, offset int) ([]model.Location, error)
	GetLocationsOfFranchiseByName(franchiseID, count, offset int, name string) ([]model.Location, error)

	ListCategories() ([]model.Category, error)

	GetProduct(int) (model.Product, error)
	ListProducts(name string, count, offset int) ([]model.Product, error)
	GetConsumer(int) (model.Consumer, error)
	GetOrder(consumerID, orderID int) (model.Order, error)
	ListOrders(id, count, offset int) ([]model.Order, error)

	Healthcheck() error
}

type ConsumerService struct {
	fh *fHandler.FranchiseHandler
	lh *lHandler.LocationHandler
	ch *cHandler.CategoryHandler
	ph *pHandler.ProductHandler
	oh *oHandler.OrderHandler
	uh *uHandler.ProfileHandler
}

// все проверки и запросы к сервису auth / user / roles производятся на уровне сервиса (т.е. бизнес-логики)

func NewConsumerService(
	fh *fHandler.FranchiseHandler,
	lh *lHandler.LocationHandler,
	ch *cHandler.CategoryHandler,
	ph *pHandler.ProductHandler,
	oh *oHandler.OrderHandler,
	uh *uHandler.ProfileHandler) Service {
	return ConsumerService{fh: fh, uh: uh, lh: lh, ch: ch, ph: ph, oh: oh}
}

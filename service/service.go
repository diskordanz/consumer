package service

import (
	cHandler "github.com/diskordanz/consumer/pkg/category/handler"
	uHandler "github.com/diskordanz/consumer/pkg/consumer/handler"
	fHandler "github.com/diskordanz/consumer/pkg/franchise/handler"
	lHandler "github.com/diskordanz/consumer/pkg/location/handler"
	oHandler "github.com/diskordanz/consumer/pkg/order/handler"
	pHandler "github.com/diskordanz/consumer/pkg/product/handler"

	"github.com/diskordanz/consumer/service/model"
)

type Service interface {
	GetFranchise(int) (model.Franchise, error)
	ListFranchises(count, offset int) ([]model.Franchise, error)

	GetLocationsOfFranchise(franchiseID, count, offset int) ([]model.Location, error)

	ListCategories() ([]model.Category, error)

	GetProduct(int) (model.Product, error)
	ListProducts(name string, count, offset int) ([]model.Product, error)
	ListProductsByCategory(id uint64, name string, count, offset int) ([]model.Product, error)

	GetOrder(id int) (model.OrderWithItems, error)
	ListOrders(id, count, offset int) ([]model.Order, error)
	CreateOrder(item model.Order) (model.Order, error)
	CreateOrderItem(item model.OrderItem) (model.OrderItem, error)

	GetConsumer(int) (model.Consumer, error)
	CreateConsumer(model.Consumer) (model.Consumer, error)
	UpdateConsumer(model.Consumer) (model.Consumer, error)
	GetCart(id, count, offset int) ([]model.CartItemsByFranchise, error)
	GetCartItem(id, productID int) (model.CartItem, error)
	CreateCartItem(model.CartItem) (model.CartItem, error)
	UpdateCartItem(model.CartItem) (model.CartItem, error)
	DeleteCartItem(model.CartItem) error

	Healthcheck() error
}

type ConsumerService struct {
	fh *fHandler.FranchiseHandler
	lh *lHandler.LocationHandler
	ch *cHandler.CategoryHandler
	ph *pHandler.ProductHandler
	oh *oHandler.OrderHandler
	uh *uHandler.ConsumerHandler
}

// все проверки и запросы к сервису auth / user / roles производятся на уровне сервиса (т.е. бизнес-логики)

func NewConsumerService(
	fh *fHandler.FranchiseHandler,
	lh *lHandler.LocationHandler,
	ch *cHandler.CategoryHandler,
	ph *pHandler.ProductHandler,
	oh *oHandler.OrderHandler,
	uh *uHandler.ConsumerHandler) Service {
	return ConsumerService{fh: fh, uh: uh, lh: lh, ch: ch, ph: ph, oh: oh}
}

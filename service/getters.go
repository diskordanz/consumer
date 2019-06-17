package service

import (
	"errors"

	pkgFranchiseModel "github.com/diskordanz/web-consumer/pkg/franchise/model"
	pkgProductModel "github.com/diskordanz/web-consumer/pkg/product/model"

	"github.com/diskordanz/web-consumer/service/model"
)

func (s ConsumerService) GetFranchise(id int) (model.Franchise, error) {
	franchise, err := s.fh.GetFranchise(id)
	if err != nil {
		return model.Franchise{}, err
	}
	resultFranchise := model.MapToFranchise(franchise)
	return resultFranchise, nil
}

func (s ConsumerService) ListFranchises(count, offset int) ([]model.Franchise, error) {
	list, err := s.fh.ListFranchises(count, offset)
	if err != nil {
		return nil, err
	}
	result := model.MapToFranchiseList(list)
	return result, nil
}

/////////////////////////////////////////////////////

func (s ConsumerService) GetLocationsOfFranchise(franchiseID, count, offset int) ([]model.Location, error) {
	locations, err := s.lh.GetLocationsOfFranchise(franchiseID, count, offset)
	if err != nil {
		return nil, err
	}
	result := model.MapToLocationsList(locations)
	return result, nil
}

/*
func (s ConsumerService) GetLocationsOfFranchiseByName(franchiseID, count, offset int, name string) ([]model.Location, error) {
	locations, err := s.lh.GetLocationsOfFranchiseByName(franchiseID, count, offset, name)
	if err != nil {
		return nil, err
	}
	result := model.MapToLocationsList(locations)
	return result, nil
}*/

/////////////////////////////////////////////////////

func (s ConsumerService) ListCategories() ([]model.Category, error) {
	list, err := s.ch.ListCategories()
	if err != nil {
		return nil, err
	}
	result := model.MapToCategoryList(list)
	return result, nil
}

/////////////////////////////////////////////////////

func (s ConsumerService) ListOrders(id, count, offset int) ([]model.Order, error) {
	list, err := s.oh.ListOrders(id, count, offset)
	if err != nil {
		return nil, err
	}
	result := model.MapToOrderList(list)
	return result, nil
}

func (s ConsumerService) GetOrder(id int) (model.OrderWithItems, error) {
	items, err := s.oh.GetOrderItems(id)
	if err != nil {
		return model.OrderWithItems{}, err
	}
	order, err := s.oh.GetOrder(id)
	if err != nil {
		return model.OrderWithItems{}, err
	}

	var products []pkgProductModel.Product
	for _, k := range items {
		product, err := s.ph.GetProductById(int(k.ProductID))
		if err != nil {
			return model.OrderWithItems{}, err
		}
		products = append(products, product)
	}

	resultOrder := model.MapToOrderWithItems(order, items, products)

	return resultOrder, nil
}

func (s ConsumerService) CreateOrderItem(item model.OrderItem) (model.OrderItem, error) {
	dbItem := model.MapToOrderItemDB(item)
	_, err := s.oh.CreateOrderItem(dbItem)
	if err != nil {
		return model.OrderItem{}, err
	}
	return item, nil
}

func (s ConsumerService) CreateOrder(item model.Order) (model.Order, error) {
	dbItem := model.MapToOrderDB(item)
	order, err := s.oh.CreateOrder(dbItem)
	if err != nil {
		return model.Order{}, err
	}
	item.ID = order.ID
	return item, nil
}

/////////////////////////////////////////////////////////////

func (s ConsumerService) GetProduct(id int) (model.Product, error) {
	product, err := s.ph.GetProductById(id)
	if err != nil {
		return model.Product{}, err
	}
	resultProduct := model.MapToProduct(product)
	return resultProduct, nil
}

func (s ConsumerService) ListProducts(name string, count, offset int) ([]model.Product, error) {
	products, err := s.ph.ListProducts(name, count, offset)
	if err != nil {
		return nil, err
	}

	var franchises []pkgFranchiseModel.Franchise
	for _, k := range products {
		franchise, err := s.fh.GetFranchise(int(k.FranchiseID))
		if err != nil {
			return nil, err
		}
		franchises = append(franchises, franchise)
	}

	result := model.MapToProductList(products, franchises)
	return result, nil
}

func (s ConsumerService) ListProductsByCategory(id uint64, name string, count, offset int) ([]model.Product, error) {
	products, err := s.ph.ListProductsByCategory(id, name, count, offset)
	if err != nil {
		return nil, err
	}
	var franchises []pkgFranchiseModel.Franchise
	for _, k := range products {
		franchise, err := s.fh.GetFranchise(int(k.FranchiseID))
		if err != nil {
			return nil, err
		}
		franchises = append(franchises, franchise)
	}

	result := model.MapToProductList(products, franchises)
	return result, nil
}

////////////////////////////////////////////////////////////////

func (s ConsumerService) GetConsumer(id int) (model.Consumer, error) {
	consumer, err := s.uh.GetConsumer(id)
	if err != nil {
		return model.Consumer{}, err
	}
	resultConsumer := model.MapToConsumer(consumer)
	return resultConsumer, nil
}

func (s ConsumerService) GetCart(id, count, offset int) ([]model.CartItemsByFranchise, error) {
	cart, err := s.uh.GetCart(id, count, offset)
	if err != nil {
		return nil, err
	}
	var products []pkgProductModel.Product
	var franchises []pkgFranchiseModel.Franchise
	for _, k := range cart {
		product, err := s.ph.GetProductById(int(k.ProductID))
		franchise, err := s.fh.GetFranchise(int(product.FranchiseID))
		if err != nil {
			return nil, err
		}
		products = append(products, product)
		franchises = append(franchises, franchise)
	}

	result := model.MapToCart(cart, products, franchises)
	return result, nil
}

func (s ConsumerService) GetCartItem(id, productID int) (model.CartItem, error) {
	cart, err := s.uh.GetCartItem(id, productID)
	if err != nil {
		return model.CartItem{
			ID:         0,
			ConsumerID: uint64(id),
			Product:    model.Product{ID: uint64(productID)},
			Count:      1,
		}, err
	}

	result := model.MapToCartItem(cart, pkgProductModel.Product{})
	return result, nil
}

func (s ConsumerService) CreateCartItem(item model.CartItem) (model.CartItem, error) {
	dbItem := model.MapToCartItemDB(item)
	_, err := s.uh.CreateCartItem(dbItem)
	if err != nil {
		return model.CartItem{}, err
	}
	return item, nil
}

func (s ConsumerService) UpdateCartItem(item model.CartItem) (model.CartItem, error) {
	dbItem := model.MapToCartItemDB(item)
	_, err := s.uh.UpdateCartItem(dbItem)
	if err != nil {
		return model.CartItem{}, err
	}
	return item, nil
}

func (s ConsumerService) DeleteCartItem(item model.CartItem) error {
	dbItem := model.MapToCartItemDB(item)
	err := s.uh.DeleteCartItem(dbItem)
	if err != nil {
		return err
	}
	return nil
}

func (s ConsumerService) CreateConsumer(consumer model.Consumer) (model.Consumer, error) {
	dbConsumer := model.MapToConsumerDB(consumer)
	c, err := s.uh.CreateConsumer(dbConsumer)
	if err != nil {
		return model.Consumer{}, err
	}
	return model.MapToConsumer(c), nil
}

func (s ConsumerService) UpdateConsumer(consumer model.Consumer) (model.Consumer, error) {
	dbConsumer := model.MapToConsumerDB(consumer)
	dbConsumer, err := s.uh.UpdateConsumer(dbConsumer)
	if err != nil {
		return model.Consumer{}, err
	}
	return model.MapToConsumer(dbConsumer), nil
}

func (s ConsumerService) Healthcheck() error {
	if s.fh == nil || s.uh == nil || s.lh == nil || s.ch == nil || s.ph == nil || s.oh == nil {
		return errors.New("error: problem with database handlers")
	}
	return nil
}

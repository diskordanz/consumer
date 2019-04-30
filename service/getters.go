package service

import "github.com/diskordanz/consumer/service/model"

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

func (s ConsumerService) SearchFranchisesByName(count, offset int, name string) ([]model.Franchise, error) {
	list, err := s.fh.SearchFranchisesByName(count, offset, name)
	if err != nil {
		return nil, err
	}
	result := model.MapToFranchiseList(list)
	return result, nil
}

/////////////////////////////////////////////////////
/////////////////////////////////////////////////////

func (s ConsumerService) GetLocationsOfFranchise(franchiseID, count, offset int) ([]model.Location, error) {
	locations, err := s.lh.GetLocationsOfFranchise(franchiseID, count, offset)
	if err != nil {
		return nil, err
	}
	result := model.MapToLocationsList(locations)
	return result, nil
}

func (s ConsumerService) GetLocationsOfFranchiseByName(franchiseID, count, offset int, name string) ([]model.Location, error) {
	locations, err := s.lh.GetLocationsOfFranchiseByName(franchiseID, count, offset, name)
	if err != nil {
		return nil, err
	}
	result := model.MapToLocationsList(locations)
	return result, nil
}

/////////////////////////////////////////////////////
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
/////////////////////////////////////////////////////

func (s ConsumerService) ListOrders(id, count, offset int) ([]model.Order, error) {
	// list, err := s.fh.ListFranchises(count, offset)
	// if err != nil {
	// 	return nil, err
	// }
	// result := model.MapToFranchiseList(list)
	return nil, nil
}

func (s ConsumerService) GetProduct(id int) (model.Product, error) {
	// product, err := s.ph.GetProductById(id)
	// if err != nil {
	// 	return model.Product{}, err
	// }
	// resultProduct := model.MapToProduct(product)
	var x model.Product
	return x, nil
}

func (s ConsumerService) ListProducts(count, offset int) ([]model.Product, error) {
	// products, err := s.ph.ListProducts(count, offset)
	// if err != nil {
	// 	return nil, err
	// }
	// result := model.MapToProductsList(products)
	return nil, nil
}

func (s ConsumerService) ListOfFranchise(franchiseID, count, offset int) ([]model.Location, error) {
	// locations, err := s.lh.GetLocationsOfFranchise(franchiseID, count, offset)
	// if err != nil {
	// 	return nil, err
	// }
	// result := model.MapToLocationsList(locations)
	return nil, nil
}

func (s ConsumerService) GetOrder(consumerID, orderID int) (model.Order, error) {
	// product, err := s.uh.GetProductById(id)
	// if err != nil {
	// 	return model.Product{}, err
	// }
	// resultProduct := model.MapToProduct(product)
	var x model.Order
	return x, nil
}

func (s ConsumerService) GetProfile(id int) (model.Profile, error) {
	// product, err := s.uh.GetProductById(id)
	// if err != nil {
	// 	return model.Product{}, err
	// }
	// resultProduct := model.MapToProduct(product)
	var x model.Profile
	return x, nil
}

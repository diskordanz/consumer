package model

import (
	pkgProductModel "github.com/diskordanz/consumer/pkg/product/model"
)

type Product struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	FranchiseID uint64  `json:"franchise_id"`
	Count       int32   `json:"count"`
	Price       float32 `json:"price"`
}

func MapToProduct(originProduct pkgProductModel.Product) Product {
	return Product{
		ID:          originProduct.ID,
		Name:        originProduct.Name,
		Description: originProduct.Description,
		Image:       originProduct.Image,
		FranchiseID: originProduct.FranchiseID,
		Count:       originProduct.Count,
		Price:       originProduct.Price,
	}
}

func MapToProductList(originList []pkgProductModel.Product) []Product {
	resultList := make([]Product, len(originList), len(originList))
	for i, v := range originList {
		resultList[i] = MapToProduct(v)
	}
	return resultList
}

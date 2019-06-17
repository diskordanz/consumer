package model

import (
	pkgFranchiseModel "github.com/diskordanz/web-consumer/pkg/franchise/model"
	pkgProductModel "github.com/diskordanz/web-consumer/pkg/product/model"
)

type Product struct {
	ID             uint64  `json:"id"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Image          string  `json:"image"`
	FranchiseID    uint64  `json:"franchise_id"`
	FranchiseImage string  `json:"franchise_image"`
	Count          uint32  `json:"count"`
	Price          float32 `json:"price"`
	CategoryID     uint64  `json:"category_id"`
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
		CategoryID:  originProduct.CategoryID,
	}
}

func MapToProductList(originList []pkgProductModel.Product, originFranchise []pkgFranchiseModel.Franchise) []Product {
	resultList := make([]Product, len(originList), len(originList))
	for i, v := range originList {
		resultList[i] = MapToProduct(v)
		resultList[i].FranchiseImage = originFranchise[i].Image
	}
	return resultList
}

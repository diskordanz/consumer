package model

import (
	pkgProductModel "github.com/diskordanz/consumer/pkg/product/model"
)

type Product struct {
	ID        uint64 `json:"id"`
	CountryID int32  `json:"country_id"`
	Name      string `json:"name"`
}

func MapToProduct(originProduct pkgProductModel.Product) Product {
	return Product{ID: originProduct.ID, Name: originProduct.Name}
}

func MapToProductList(originList []pkgProductModel.Product) []Product {
	resultList := make([]Product, len(originList), len(originList))
	for i, v := range originList {
		resultList[i] = MapToProduct(v)
	}
	return resultList
}

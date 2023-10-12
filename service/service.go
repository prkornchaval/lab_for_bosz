package service

import (
	"fmt"
	"labForBosz/domain"
)

type ProductServiceInterface interface {
	GetProductList() []domain.Product
	GetProductById(id string) *domain.Product
}

var productList = []domain.Product{
	{Id: "1", Name: "Cola", Price: 14},
	{Id: "2", Name: "Pen", Price: 5},
	{Id: "3", Name: "Book", Price: 10},
}

type productService struct{}

func New() ProductServiceInterface {
	return &productService{}
}

func (p *productService) GetProductList() []domain.Product {
	return productList
}

func (p *productService) GetProductById(id string) *domain.Product {
	// for i := 0; i < len(productList); i++ {
	// 	if productList[i].Id == id {
	// 		productList[i].Id = id
	// 		return &productList[i]
	// 	}
	// }
	for k, product := range productList {
		fmt.Println(k)
		if product.Id == id {
			return &product
		}
	}
	return nil
}

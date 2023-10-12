package service_test

import (
	"labForBosz/domain"
	"labForBosz/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceProduct(t *testing.T) {
	t.Run("1_get_product_list", func(t *testing.T) {
		svc := service.New()
		res := svc.GetProductList()
		assert.Equal(t, []domain.Product{
			{Id: "1", Name: "Cola", Price: 14},
			{Id: "2", Name: "Pen", Price: 5},
			{Id: "3", Name: "Book", Price: 10},
		}, res)
	})

	t.Run("2_get_product_by_id", func(t *testing.T) {
		svc := service.New()
		id := "2"
		res := svc.GetProductById(id)
		assert.Equal(t, &domain.Product{
			Id: "2", Name: "Pen", Price: 5,
		}, res)
	})

	t.Run("3_get_product_by_id", func(t *testing.T) {
		svc := service.New()
		id := "5"
		res := svc.GetProductById(id)
		assert.Equal(t, (*domain.Product)(nil), res)
	})

}

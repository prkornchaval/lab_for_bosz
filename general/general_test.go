package general_test

import (
	"labForBosz/domain"
	"labForBosz/general"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomer(t *testing.T) {
	t.Run("1_TestNewCustomer", func(t *testing.T) {
		name := "test"
		age := 28
		res := general.NewCustomer(name, age)
		assert.Equal(t, domain.Customer{Name: "test", Age: 28}, res)
	})

	t.Run("2_TestChangeCustomerData", func(t *testing.T) {
		cus := domain.Customer{
			Name: "test",
			Age:  26,
		}
		name := "Coca"
		age := 30

		general.ChangeCustomerData(&cus, name, age)
		assert.Equal(t, domain.Customer{Name: "Coca", Age: 30}, cus)
	})

}

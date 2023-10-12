package general

import (
	"fmt"
	"labForBosz/domain"
)

func NewCustomer(name string, age int) domain.Customer {
	return domain.Customer{Name: name, Age: age}
}

func ChangeCustomerData(cus *domain.Customer, name string, age int) {
	cus.Name = name
	cus.Age = age
}

func X() {
	res := NewCustomer("fff", 26)
	fmt.Printf("obj: %+v", res)
}

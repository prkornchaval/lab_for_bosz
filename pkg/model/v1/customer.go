package model

type CreateCustomerRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

type GetCustomerReqeust struct {
	Id int `uri:"id"`
}

type GetCustomerResponse struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

type GetCustomerNameResponse struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
}

func (g *GetCustomerReqeust) XX() {

}

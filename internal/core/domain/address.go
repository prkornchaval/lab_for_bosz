package domain

type Address struct {
	Id          int    `json:"id" db:"id"`
	Owner       int    `json:"owner" db:"owner"`
	Subdistrict string `json:"subdistrict" db:"subdistrict"`
	District    string `json:"district" db:"district"`
	Province    string `json:"province" db:"province"`
}

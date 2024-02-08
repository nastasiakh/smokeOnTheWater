package models

type Address struct {
	Country   string `json:"country"`
	Region    string `json:"region"`
	Location  string `json:"location"`
	Street    string `json:"street"`
	HouseNum  string `json:"houseNum"`
	Apartment uint   `json:"apartment"`
	Zipcode   string `json:"zipcode"`
}

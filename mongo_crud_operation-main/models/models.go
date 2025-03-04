package models

type User struct {
	Name    string  `json:"name" bson:"user_name"`
	Age     int     `json:"age" bson:"user_age"`
	Address Address `json:"address" bson:"user_address"`
	Gender  string `json:"gender" bson:"user_gender"`
	Company string `json:"company" bson:"user_company"`
}

type Address struct {
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
	Pincode int    `json:"pincode" bson:"pincode"`
	Nation string `json:"nation" bson:"nation"`
}

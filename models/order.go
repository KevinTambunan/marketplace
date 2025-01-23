package models

import "time"

type Order struct {
	Id          int       `json:"id"`
	User_id     int       `json:"user_id"`
	Product_id  int       `json:"product_id"`
	Quantity    int       `json:"quantity"`
	Total_price int       `json:"total_price"`
	Status_id   int       `json:"status_id"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

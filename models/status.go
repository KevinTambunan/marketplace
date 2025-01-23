package models

import "time"

type Status struct {
	Id         int       `json:"id"`
	Name       string    `json:"Name"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

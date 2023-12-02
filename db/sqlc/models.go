package db

import (
	"time"
)

//Car model
type Car struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Price     int32     `json:"price"`
	Brand     string    `json:"brand"`
	CreatedAt time.Time `json:"created_at"`
}

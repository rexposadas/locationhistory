package models

import "time"

// Order represents an order in our system.
type Order struct {
	ID      string     `json:"order_id"`
	History []Location `json:"history"`
}

type Location struct {
	Expiry *time.Time `json:"expiry"`

	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

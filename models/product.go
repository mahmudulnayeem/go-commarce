package models

import "time"

type Product struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"imageUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}
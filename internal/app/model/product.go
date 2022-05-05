package model

type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Price       string `json:"price"`
	Description string `json:"description"`
}

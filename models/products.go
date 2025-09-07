package models

type Product struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Price   string `json:"price"`
	Details string `json:"details"`
}

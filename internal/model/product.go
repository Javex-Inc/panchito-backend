package model

type Product struct {
	ID          string
	Name        string
	Description string
	Image       string
	Price       float32
	Category    string // TODO: Will be Category type
	IsOnSale    bool   // When the product is an offer
}

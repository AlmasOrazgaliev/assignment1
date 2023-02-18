package model

type Item struct {
	Id          int
	Name        string
	Price       int
	Description string
	Sold        float64
	Rating      float64
	Moderated   bool
	Seller_id   int
}

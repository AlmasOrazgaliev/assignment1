package model

type Item struct {
	Id          int
	Name        string
	Price       int
	Description string
	Sold        int
	Rating      int
	Moderated   bool
	Seller_id   int
}

package models

type Item struct {
	ID          int
	Name        string
	Price       int
	Description string
	Quantity    int
	Category    string
	Subcategory string
	Slug        string
	Store       string
	Link        string
	Images      []string
}

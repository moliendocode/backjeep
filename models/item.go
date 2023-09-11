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

type CreateItemRequest struct {
	Name        string   `json:"name"`
	Price       int      `json:"price"`
	Description string   `json:"description"`
	Quantity    int      `json:"quantity"`
	Category    string   `json:"category"`
	Subcategory string   `json:"subcategory"`
	Slug        string   `json:"slug"`
	Store       string   `json:"store"`
	Link        string   `json:"link"`
	Images      []string `json:"images"`
}

type UpdateItemRequest struct {
	CreateItemRequest
	ImagesToDelete []string `json:"imagesToDelete"`
	NewImages      []string `json:"newImages"`
}

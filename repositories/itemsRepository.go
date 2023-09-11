package repositories

import (
	"backjeep/models"
	"backjeep/utils"
	"strconv"
)

type ItemRepository interface {
	GetAllItems() ([]models.Item, error)
	CreateItem(req models.CreateItemRequest) (models.Item, error)
	GetItemDetails(itemID string) (models.Item, error)
	UpdateItem(itemID int, req models.UpdateItemRequest) (models.Item, error)
}

type ItemRepo struct{}

func (ir *ItemRepo) GetAllItems() ([]models.Item, error) {
	var items []models.Item

	rows, err := utils.DB.Query(`SELECT id, name, price, description, quantity, category, subcategory, slug, store, link FROM items`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		err = rows.Scan(&item.ID, &item.Name, &item.Price, &item.Description, &item.Quantity, &item.Category, &item.Subcategory, &item.Slug, &item.Store, &item.Link)
		if err != nil {
			return nil, err
		}

		imgRows, err := utils.DB.Query(`SELECT url FROM item_images WHERE item_id=$1`, item.ID)
		if err != nil {
			return nil, err
		}
		for imgRows.Next() {
			var imageURL string
			err = imgRows.Scan(&imageURL)
			if err != nil {
				imgRows.Close()
				return nil, err
			}
			item.Images = append(item.Images, imageURL)
		}
		imgRows.Close()

		items = append(items, item)
	}

	return items, nil
}

func (ir *ItemRepo) CreateItem(req models.CreateItemRequest) (models.Item, error) {
	query := `
		INSERT INTO items (name, price, description, quantity, category, subcategory, slug, store, link)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
		RETURNING id
	`
	var itemID int
	err := utils.DB.QueryRow(query, req.Name, req.Price, req.Description, req.Quantity, req.Category, req.Subcategory, req.Slug, req.Store, req.Link).Scan(&itemID)
	if err != nil {
		return models.Item{}, err
	}

	for _, imageURL := range req.Images {
		_, err = utils.DB.Exec(`INSERT INTO item_images (item_id, url) VALUES ($1, $2)`, itemID, imageURL)
		if err != nil {
			return models.Item{}, err
		}
	}

	return models.Item{
		ID:          itemID,
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Quantity:    req.Quantity,
		Category:    req.Category,
		Subcategory: req.Subcategory,
		Slug:        req.Slug,
		Store:       req.Store,
		Link:        req.Link,
		Images:      req.Images,
	}, nil
}

func (ir *ItemRepo) GetItemDetails(itemID string) (models.Item, error) {
	var item models.Item

	err := utils.DB.QueryRow(`SELECT id, name, price, description, quantity, category, subcategory, slug, store, link FROM items WHERE id = $1`, itemID).
		Scan(&item.ID, &item.Name, &item.Price, &item.Description, &item.Quantity, &item.Category, &item.Subcategory, &item.Slug, &item.Store, &item.Link)
	if err != nil {
		return models.Item{}, err
	}

	imgRows, err := utils.DB.Query(`SELECT url FROM item_images WHERE item_id=$1`, item.ID)
	if err != nil {
		return models.Item{}, err
	}
	defer imgRows.Close()
	for imgRows.Next() {
		var imageURL string
		err = imgRows.Scan(&imageURL)
		if err != nil {
			return models.Item{}, err
		}
		item.Images = append(item.Images, imageURL)
	}

	return item, nil
}

func (ir *ItemRepo) UpdateItem(itemID int, req models.UpdateItemRequest) (models.Item, error) {
	query := `
        UPDATE items 
        SET name=$2, price=$3, description=$4, quantity=$5, category=$6, subcategory=$7, slug=$8, store=$9, link=$10
        WHERE id=$1
    `
	_, err := utils.DB.Exec(query, itemID, req.Name, req.Price, req.Description, req.Quantity, req.Category, req.Subcategory, req.Slug, req.Store, req.Link)
	if err != nil {
		return models.Item{}, err
	}

	for _, imageURL := range req.ImagesToDelete {
		_, err = utils.DB.Exec(`DELETE FROM item_images WHERE item_id=$1 AND url=$2`, itemID, imageURL)
		if err != nil {
			return models.Item{}, err
		}
	}

	for _, imageURL := range req.NewImages {
		_, err = utils.DB.Exec(`INSERT INTO item_images (item_id, url) VALUES ($1, $2)`, itemID, imageURL)
		if err != nil {
			return models.Item{}, err
		}
	}

	updatedItem, err := ir.GetItemDetails(strconv.Itoa(itemID))
	if err != nil {
		return models.Item{}, err
	}

	return updatedItem, nil
}

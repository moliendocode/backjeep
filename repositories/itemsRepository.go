package repositories

import (
	"backjeep/models"
	"backjeep/utils"
)

type ItemRepository interface {
	GetAllItems() ([]models.Item, error)
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

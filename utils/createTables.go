package utils

func CreateTable() error {
	// Users table
	usersTable := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL,
            password VARCHAR(100) NOT NULL
        );
    `
	_, err := DB.Exec(usersTable)
	if err != nil {
		return err
	}

	// Items table
	itemsTable := `
        CREATE TABLE IF NOT EXISTS items (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            price INTEGER NOT NULL CHECK (price >= 0),
            description TEXT NOT NULL,
            quantity INT DEFAULT 0 NOT NULL,
            category VARCHAR(100) NOT NULL,
            subcategory VARCHAR(100) NOT NULL,
            slug VARCHAR(255) UNIQUE NOT NULL,
            store VARCHAR(255) NOT NULL,
            link TEXT NOT NULL
        );
    `
	_, err = DB.Exec(itemsTable)
	if err != nil {
		return err
	}

	// Item images table
	itemImagesTable := `
        CREATE TABLE IF NOT EXISTS item_images (
            id SERIAL PRIMARY KEY,
            item_id INTEGER REFERENCES items(id) ON DELETE CASCADE,
            url TEXT NOT NULL
        );
    `
	_, err = DB.Exec(itemImagesTable)
	if err != nil {
		return err
	}
	return nil
}

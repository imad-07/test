package data

import (
	"database/sql"
)

type InfoData struct {
	Db *sql.DB
}

// Get Categories
func (inf *InfoData) GetCategories() ([]string, error) {
	rows, err := inf.Db.Query("SELECT category_name FROM categories")
	if err != nil {
		return nil, err
	}

	var categories []string
	for rows.Next() {

		var category string
		err = rows.Scan(&category)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

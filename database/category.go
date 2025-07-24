package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db   *sql.DB
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCategory(db *sql.DB, id, name string) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name) VALUES ($1, $2)", id, name)
	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name}, nil
}

func (c *Category) GetAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

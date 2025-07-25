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

func NewCategory(db *sql.DB) *Category {
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

func (c *Category) GetByProductID(id string) (Category, error) {
	var category Category
	err := c.db.QueryRow("SELECT c.id, c.name FROM categories c JOIN products p ON c.id = p.category_id WHERE p.id = $1", id).Scan(&category.ID, &category.Name)
	if err != nil {
		return Category{}, err
	}
	return category, nil
}

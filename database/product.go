package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Product struct {
	db          *sql.DB
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryID  string `json:"category_id,omitempty"`
}

func NewProduct(db *sql.DB) *Product {
	return &Product{db: db}
}

func (p *Product) Create(name, description, categoryID string) (Product, error) {
	id := uuid.New().String()
	_, err := p.db.Exec("INSERT INTO products (id, name, description, category_id) VALUES ($1, $2, $3, $4)", id, name, description, categoryID)
	if err != nil {
		return Product{}, err
	}

	return Product{ID: id, Name: name, Description: description, CategoryID: categoryID}, nil
}

func (p *Product) GetAll() ([]Product, error) {
	rows, err := p.db.Query("SELECT id, name, description, category_id FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (p *Product) GetAllByCategoryID(id string) ([]Product, error) {
	rows, err := p.db.Query("SELECT id, name, description FROM products WHERE category_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, sql.ErrNoRows
	}

	return products, nil
}

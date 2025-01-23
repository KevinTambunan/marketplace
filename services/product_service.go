package services

import (
	"database/sql"
	"errors"
	"marketplace/models"
	"time"
)

type ProductService struct {
	DB *sql.DB
}

func NewProductService(db *sql.DB) *ProductService {
	return &ProductService{DB: db}
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	query := "SELECT id, name, description, price, category_id, stock, created_at, updated_at FROM products"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Category_id, &product.Stock, &product.Created_at, &product.Updated_at); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (s *ProductService) CreateProduct(name, description string, price, categoryId, stock int) (int, error) {
	if name == "" || description == "" || price <= 0 || categoryId <= 0 || stock < 0 {
		return 0, errors.New("invalid product input")
	}

	query := "INSERT INTO products (name, description, price, category_id, stock, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := s.DB.Exec(query, name, description, price, categoryId, stock, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	productID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(productID), nil
}

func (s *ProductService) GetProductById(id int) (*models.Product, error) {
	query := "SELECT id, name, description, price, category_id, stock, created_at, updated_at FROM products WHERE id = ?"
	row := s.DB.QueryRow(query, id)

	var product models.Product
	if err := row.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Category_id, &product.Stock, &product.Created_at, &product.Updated_at); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &product, nil
}

func (s *ProductService) UpdateProduct(id int, name, description string, price, categoryId, stock int) error {
	if name == "" || description == "" || price <= 0 || categoryId <= 0 || stock < 0 {
		return errors.New("invalid product input")
	}

	query := "UPDATE products SET name = ?, description = ?, price = ?, category_id = ?, stock = ?, updated_at = ? WHERE id = ?"
	_, err := s.DB.Exec(query, name, description, price, categoryId, stock, time.Now(), id)
	return err
}

func (s *ProductService) DeleteProduct(id int) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := s.DB.Exec(query, id)
	return err
}

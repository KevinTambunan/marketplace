package services

import (
	"database/sql"
	"errors"
	"marketplace/models"
	"time"
)

type CategoryService struct {
	DB *sql.DB
}

func NewCategoryService(db *sql.DB) *CategoryService {
	return &CategoryService{DB: db}
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	query := "SELECT id, name, created_at, updated_at FROM categories"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.Created_at, &category.Updated_at); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (s *CategoryService) CreateCategory(name string) error {
	if name == "" {
		return errors.New("category name cannot be empty")
	}

	query := "INSERT INTO categories (name, created_at, updated_at) VALUES (?, ?, ?)"
	_, err := s.DB.Exec(query, name, time.Now(), time.Now())
	return err
}

func (s *CategoryService) GetCategoryById(id int) (*models.Category, error) {
	query := "SELECT id, name, created_at, updated_at FROM categories WHERE id = ?"
	row := s.DB.QueryRow(query, id)

	var category models.Category
	if err := row.Scan(&category.Id, &category.Name, &category.Created_at, &category.Updated_at); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	return &category, nil
}

func (s *CategoryService) UpdateCategory(id int, name string) error {
	if name == "" {
		return errors.New("category name cannot be empty")
	}

	query := "UPDATE categories SET name = ?, updated_at = ? WHERE id = ?"
	_, err := s.DB.Exec(query, name, time.Now(), id)
	return err
}

func (s *CategoryService) DeleteCategory(id int) error {
	query := "DELETE FROM categories WHERE id = ?"
	_, err := s.DB.Exec(query, id)
	return err
}

package services

import (
	"database/sql"
	"errors"
	"marketplace/models"
	"time"
)

type StatusService struct {
	DB *sql.DB
}

func NewStatusService(db *sql.DB) *StatusService {
	return &StatusService{DB: db}
}

func (s *StatusService) GetAllStatuses() ([]models.Status, error) {
	query := "SELECT id, name, created_at, updated_at FROM status"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []models.Status
	for rows.Next() {
		var status models.Status
		if err := rows.Scan(&status.Id, &status.Name, &status.Created_at, &status.Updated_at); err != nil {
			return nil, err
		}
		statuses = append(statuses, status)
	}
	return statuses, nil
}

func (s *StatusService) CreateStatus(name string) (int, error) {
	if name == "" {
		return 0, errors.New("status name cannot be empty")
	}

	query := "INSERT INTO status (name, created_at, updated_at) VALUES (?, ?, ?)"
	result, err := s.DB.Exec(query, name, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	statusID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(statusID), nil
}

func (s *StatusService) GetStatusById(id int) (*models.Status, error) {
	query := "SELECT id, name, created_at, updated_at FROM status WHERE id = ?"
	row := s.DB.QueryRow(query, id)

	var status models.Status
	if err := row.Scan(&status.Id, &status.Name, &status.Created_at, &status.Updated_at); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("status not found")
		}
		return nil, err
	}

	return &status, nil
}

func (s *StatusService) UpdateStatus(id int, name string) error {
	if name == "" {
		return errors.New("status name cannot be empty")
	}

	query := "UPDATE status SET name = ?, updated_at = ? WHERE id = ?"
	_, err := s.DB.Exec(query, name, time.Now(), id)
	return err
}

func (s *StatusService) DeleteStatus(id int) error {
	query := "DELETE FROM status WHERE id = ?"
	_, err := s.DB.Exec(query, id)
	return err
}

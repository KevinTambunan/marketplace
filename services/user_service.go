package services

import (
	"database/sql"
	"errors"
	"log"
	"marketplace/models"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) GetUserByID(userID int) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE id = ?`
	err := s.DB.QueryRow(query, userID).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.Created_at, &user.Updated_at)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		log.Println(err)
		return nil, err
	}

	return user, nil
}

func (s *UserService) Register(name, email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	query := `INSERT INTO users (name, email, password, role, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())`
	result, err := s.DB.Exec(query, name, email, hashedPassword, "user")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &models.User{
		Id:       int(userID),
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     "user",
	}, nil
}

func (s *UserService) Login(email, password string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE email = ?`
	err := s.DB.QueryRow(query, email).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.Created_at, &user.Updated_at)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid credentials")
		}
		log.Println(err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

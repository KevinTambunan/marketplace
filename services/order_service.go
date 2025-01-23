package services

import (
	"database/sql"
	"errors"
	"marketplace/models"
	"time"
)

type OrderService struct {
	DB *sql.DB
}

func NewOrderService(db *sql.DB) *OrderService {
	return &OrderService{DB: db}
}

func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	query := "SELECT id, user_id, product_id, quantity, total_price, status_id, created_at, updated_at FROM orders"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.Id, &order.User_id, &order.Product_id, &order.Quantity, &order.Total_price, &order.Status_id, &order.Created_at, &order.Updated_at); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (s *OrderService) CreateOrder(userId, productId, quantity, statusId int) (int, error) {
	if userId <= 0 || productId <= 0 || quantity <= 0 || statusId <= 0 {
		return 0, errors.New("invalid order input")
	}

	var price int
	query := "SELECT price FROM products WHERE id = ?"
	err := s.DB.QueryRow(query, productId).Scan(&price)
	if err != nil {
		return 0, err
	}

	totalPrice := price * quantity

	query = "INSERT INTO orders (user_id, product_id, quantity, total_price, status_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := s.DB.Exec(query, userId, productId, quantity, totalPrice, statusId, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(orderID), nil
}

func (s *OrderService) GetOrderById(id int) (*models.Order, error) {
	query := "SELECT id, user_id, product_id, quantity, total_price, status_id, created_at, updated_at FROM orders WHERE id = ?"
	row := s.DB.QueryRow(query, id)

	var order models.Order
	if err := row.Scan(&order.Id, &order.User_id, &order.Product_id, &order.Quantity, &order.Total_price, &order.Status_id, &order.Created_at, &order.Updated_at); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("order not found")
		}
		return nil, err
	}

	return &order, nil
}

func (s *OrderService) UpdateOrder(id, userId, productId, quantity, statusId int) error {
	if userId <= 0 || productId <= 0 || quantity <= 0 || statusId <= 0 {
		return errors.New("invalid order input")
	}

	var price int
	query := "SELECT price FROM products WHERE id = ?"
	err := s.DB.QueryRow(query, productId).Scan(&price)
	if err != nil {
		return err
	}

	totalPrice := price * quantity

	query = "UPDATE orders SET user_id = ?, product_id = ?, quantity = ?, total_price = ?, status_id = ?, updated_at = ? WHERE id = ?"
	_, err = s.DB.Exec(query, userId, productId, quantity, totalPrice, statusId, time.Now(), id)
	return err
}

func (s *OrderService) DeleteOrder(id int) error {
	query := "DELETE FROM orders WHERE id = ?"
	_, err := s.DB.Exec(query, id)
	return err
}

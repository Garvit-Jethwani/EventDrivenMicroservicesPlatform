package database

import (
	"database/sql"
	"errors"

	"github.com/Garvit-Jethwani/order-serice/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDatabase(dataSourceName string) error {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	return db.Ping()
}

func CreateOrder(order *models.Order) error {
	order.ID = uuid.New().String()
	order.Status = "created"

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO orders (id, customer_id, status) VALUES ($1, $2, $3)",
		order.ID, order.CustomerID, order.Status)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range order.Items {
		_, err = tx.Exec("INSERT INTO order_items (order_id, product_id, quantity) VALUES ($1, $2, $3)",
			order.ID, item.ProductID, item.Quantity)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func GetOrderById(orderId string) (*models.Order, error) {
	order := &models.Order{}

	err := db.QueryRow("SELECT id, customer_id, status FROM orders WHERE id = $1", orderId).
		Scan(&order.ID, &order.CustomerID, &order.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("order not found")
		}
		return nil, err
	}

	rows, err := db.Query("SELECT product_id, quantity FROM order_items WHERE order_id = $1", orderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := models.OrderItem{}
		if err := rows.Scan(&item.ProductID, &item.Quantity); err != nil {
			return nil, err
		}
		order.Items = append(order.Items, item)
	}

	return order, nil
}

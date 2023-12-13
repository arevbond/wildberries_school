package postgres

import (
	"encoding/json"
	"first_task_l0/internal/config"
	"first_task_l0/internal/storage"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Storage struct {
	db *sqlx.DB
}

func New(cfg *config.Config) (*Storage, error) {
	dbSource := fmt.Sprintf("postgres://%s:%s@localhost:5431/%s", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDBName)
	conn, err := sqlx.Connect("pgx", dbSource)
	if err != nil {
		return nil, errors.Wrap(err, "sqlx connect")
	}

	err = conn.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "ping failed")
	}
	return &Storage{db: conn}, nil
}

// CreateCustomer записывает данные покупателя в базу данных.
func (s *Storage) CreateCustomer(customer *storage.CustomerDB) error {
	q := `INSERT INTO customers (customer_id, name, phone, zip, city, address, region, email)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := s.db.Exec(q, customer.CustomerID, customer.Name, customer.Phone, customer.Zip,
		customer.City, customer.Address, customer.Region, customer.Email)
	if err != nil {
		return errors.Wrap(err, "can't insert customer in db:")
	}
	return nil
}

// CreateOrder записывает данные о заказе в базу данных.
func (s *Storage) CreateOrder(order *storage.OrderDB) error {
	q := `INSERT INTO orders (order_uid, track_number, entry, customer_id, payment,
            locale, internal_signature, delivery_service, shard_key, sm_id, date_created, oof_shard)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err := s.db.Exec(q, order.OrderUID, order.TrackNumber, order.Entry, order.CustomerID, order.PaymentDB,
		order.Locale, order.InternalSignature, order.DeliveryService, order.ShardKey, order.SmID,
		order.DateCreated, order.OofShard)
	if err != nil {
		return errors.Wrap(err, "can't insert order in db:")
	}
	return nil
}

// CreateItem создаёт запись об item в бд.
func (s *Storage) CreateItem(item *storage.ItemDB) error {
	q := `INSERT INTO items (chrt_id, track_number, price, rid, name, sale, total_price, 
                   nm_id, brand, status, order_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := s.db.Exec(q, item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name,
		item.Sale, item.TotalPrice, item.NmID, item.Brand, item.Sale, item.OrderID)
	if err != nil {
		return errors.Wrap(err, "can't insert item in db:")
	}
	return nil
}

// GetCustomer возвращает покупателя из бд по его id.
func (s *Storage) GetCustomer(id string) (*storage.CustomerDB, error) {
	q := `SELECT * from customers WHERE customer_id = $1`
	customer := storage.CustomerDB{}
	err := s.db.Get(&customer, q, id)
	if err != nil {
		return nil, errors.Wrap(err, "can't get customer by id")
	}
	return &customer, nil
}

func (s *Storage) GetItems(orderID string) ([]storage.ItemDB, error) {
	q := `SELECT * from items WHERE order_id = $1`
	items := make([]storage.ItemDB, 0)
	err := s.db.Select(&items, q, orderID)
	if err != nil {
		return nil, errors.Wrap(err, "can't get items by order_id")
	}
	return items, nil
}

// GetOrders возращает массив JSONов.
func (s *Storage) GetOrders(limit int) ([]storage.OrderDB, error) {
	rows, err := s.db.Query("SELECT * from orders LIMIT $1;", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []storage.OrderDB{}
	for rows.Next() {
		var order storage.OrderDB
		var data []byte
		if err := rows.Scan(&order.OrderUID, &order.TrackNumber, &order.Entry, &order.CustomerID, &data,
			&order.Locale, &order.InternalSignature, &order.DeliveryService, &order.ShardKey, &order.SmID, &order.DateCreated,
			&order.OofShard); err != nil {
			return orders, err
		}
		err := json.Unmarshal(data, &order.PaymentDB)
		if err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return orders, err
	}
	return orders, nil
}

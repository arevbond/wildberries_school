package postgres

import (
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
	q := `INSERT INTO customers (id, name, phone, zip, city, address, region, email)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := s.db.Exec(q, customer.ID, customer.Name, customer.Phone, customer.Zip,
		customer.City, customer.Address, customer.Region, customer.Email)
	if err != nil {
		return errors.Wrap(err, "can't insert customer in db:")
	}
	return nil
}

// CreateOrder записывает данные о заказе в базу данных.
func (s *Storage) CreateOrder(order *storage.OrderDB) error {
	q := `INSERT INTO orders (id, track_number, entry, customer_id, payment,
            locale, internal_signature, delivery_service, shard_key, sm_id, date_created, oof_shard)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err := s.db.Exec(q, order.ID, order.TrackNumber, order.Entry, order.CustomerID, order.PaymentDB,
		order.Locale, order.InternalSignature, order.DeliveryService, order.ShardKey, order.SmID,
		order.DateCreated, order.OofShard)
	if err != nil {
		return errors.Wrap(err, "can't insert order in db:")
	}
	return nil
}

func (s *Storage) CreateItem(item *storage.ItemDB) error {
	q := `INSERT INTO items (id, track_number, price, rid, name, sale, total_price, 
                   nm_id, brand, status, order_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := s.db.Exec(q, item.ID, item.TrackNumber, item.Price, item.Rid, item.Name,
		item.Sale, item.TotalPrice, item.NmID, item.Brand, item.Sale, item.OrderID)
	if err != nil {
		return errors.Wrap(err, "can't insert item in db:")
	}
	return nil
}

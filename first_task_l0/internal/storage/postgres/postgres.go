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

func (s *Storage) CreateCustomer(customer *storage.CustomerDB) error {
	q := `INSERT INTO customers (id, name, phone, zip, city, address, region, email)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := s.db.Exec(q, customer.ID, customer.Name, customer.Phone, customer.Zip, customer.Zip,
		customer.City, customer.Address, customer.Region, customer.Email)
	if err != nil {
		return errors.Wrap(err, "can't insert customer in db")
	}
	return nil
}

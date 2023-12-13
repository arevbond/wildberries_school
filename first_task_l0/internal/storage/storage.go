package storage

import (
	"time"
)

type Storage interface {
	CreateCustomer(customer *CustomerDB) error
	CreateOrder(order *OrderDB) error
	CreateItem(item *ItemDB) error

	GetOrders(limit int) ([]OrderDB, error)
	GetItems(orderID string) ([]ItemDB, error)
	GetCustomer(id string) (*CustomerDB, error)
}

type ItemDB struct {
	ChrtID      int    `db:"chrt_id"`
	TrackNumber string `db:"track_number"`
	Price       int    `db:"price"`
	Rid         string `db:"rid"`
	Name        string `db:"name"`
	Sale        int    `db:"sale"`
	Size        string `db:"size"`
	TotalPrice  int    `db:"total_price"`
	NmID        int    `db:"nm_id"`
	Brand       string `db:"brand"`
	Status      int    `db:"status"`
	OrderID     string `db:"order_id"`
}

type CustomerDB struct {
	CustomerID string `db:"customer_id"`
	Name       string `db:"name"`
	Phone      string `db:"phone"`
	Zip        string `db:"zip"`
	City       string `db:"city"`
	Address    string `db:"address"`
	Region     string `db:"region"`
	Email      string `db:"email"`
}

type OrderDB struct {
	OrderUID          string    `db:"order_uid"`
	TrackNumber       string    `db:"track_number"`
	Entry             string    `db:"entry"`
	CustomerID        string    `db:"customer_id"`
	PaymentDB         PaymentDB `json:"payment" db:"payment"`
	Locale            string    `db:"locale"`
	InternalSignature string    `db:"internal_signature"`
	DeliveryService   string    `db:"delivery_service"`
	ShardKey          string    `db:"shard_key"`
	SmID              int       `db:"sm_id"`
	DateCreated       time.Time `db:"date_created"`
	OofShard          string    `db:"oof_shard"`
}

type PaymentDB struct {
	Transaction  string `db:"transaction" json:"transaction"`
	RequestID    string `db:"request_id" json:"request_id"`
	Currency     string `db:"currency" json:"currency"`
	Provider     string `db:"provider" json:"provider"`
	Amount       int    `db:"amount" json:"amount"`
	PaymentDT    int    `db:"payment_dt" json:"payment_dt"`
	Bank         string `db:"bank" json:"bank"`
	DeliveryCost int    `db:"delivery_cost" json:"delivery_cost"`
	GoodsTotal   int    `db:"goods_total" json:"goods_total"`
	CustomFee    int    `db:"custom_fee" json:"custom_fee"`
}

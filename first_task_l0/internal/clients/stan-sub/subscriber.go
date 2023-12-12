package stan_sub

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log/slog"
	"time"
)

type Order struct {
	OrderUID          string `json:"order_uid" db:"order_uid"`
	TrackNumber       string `json:"track_number" db:"track_number"`
	Entry             string `json:"entry" db:"entry"`
	Delivery          `json:"delivery" db:"delivery"`
	Payment           `json:"payment" db:"payment"`
	Items             []Item    `json:"items" db:"items"`
	Locale            string    `json:"locale" db:"locale"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature"`
	CustomerID        string    `json:"customer_id" db:"customer_id"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service"`
	ShardKey          string    `json:"shard_key" db:"shard_key"`
	SmID              int       `json:"sm_id" db:"sm_id"`
	DateCreated       time.Time `json:"date_created" db:"date_created"`
	OofShard          string    `json:"oof_shard" db:"oof_shard"`
}

type Delivery struct {
	Name    string `son:"name" db:"name"`
	Phone   string `json:"phone" db:"phone"`
	Zip     string `json:"zip" db:"zip"`
	City    string `json:"city"`
	Address string `json:"address" db:"address"`
	Region  string `json:"region" db:"region"`
	Email   string `json:"email" db:"email"`
}

type Payment struct {
	Transaction  string `json:"transaction" db:"transaction"`
	RequestID    string `json:"request_id" db:"request_id"`
	Currency     string `json:"currency" db:"currency"`
	Provider     string `json:"provider" db:"provider"`
	Amount       int    `json:"amount" db:"amount"`
	PaymentID    int    `json:"paymentID" db:"paymentID"`
	Bank         string `json:"bank" db:"bank"`
	DeliveryCost int    `json:"delivery_cost" db:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total" db:"goods_total"`
	CustomFee    int    `json:"custom_fee" db:"custom_fee"`
}

type Item struct {
	ChrtID      int    `json:"chrt_id" db:"chrt_id"`
	TrackNumber string `json:"track_number" db:"track_number"`
	Price       int    `json:"price" db:"price"`
	Rid         string `json:"rid" db:"rid"`
	Name        string `json:"name" db:"name"`
	Sale        int    `json:"sale" db:"sale"`
	Size        string `json:"size" db:"size"`
	TotalPrice  int    `json:"total_price" db:"total_price"`
	NmID        int    `json:"nm_id" db:"nm_id"`
	Brand       string `json:"brand" db:"brand"`
	Status      int    `json:"status" db:"status"`
}

type Subscriber struct {
	conn   stan.Conn
	sub    stan.Subscription
	log    *slog.Logger
	Orders chan Order
}

func New(clusterID, clientID string, log *slog.Logger) *Subscriber {
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Error("can't create new Subscriber:", err)
		return nil
	}
	return &Subscriber{conn: sc, log: log, Orders: make(chan Order)}
}

func (s *Subscriber) Subscribe(subject string) error {
	sub, err := s.conn.Subscribe(subject, func(msg *stan.Msg) {
		var order Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			s.log.Error("cant unmarshall order:", err)
			return
		}
		s.log.Info(fmt.Sprintf("recieved order from nats: %s", order.Name))
		s.Orders <- order
	})
	if err != nil {
		s.log.Error(fmt.Sprintf("can't subscribe to topic %s:", subject), err)
	}
	s.sub = sub
	return nil
}

func (s *Subscriber) Unsubscribe() error {
	close(s.Orders)

	err := s.sub.Unsubscribe()
	if err != nil {
		s.log.Error("can't unsubscribe from topic:", err)
		return err
	}
	return nil
}

func (s *Subscriber) Close() error {
	err := s.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

type Order struct {
	OrderUID          string `json:"order_uid"`
	TrackNumber       string `json:"track_number"`
	Entry             string `json:"entry"`
	Delivery          `json:"delivery"`
	Payment           `json:"payment"`
	Items             []Item    `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	ShardKey          string    `json:"shard_key"`
	SmID              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

type Delivery struct {
	Name    string `son:"name"`
	Phone   string `json:"phone"`
	Zip     string `son:"zip"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentID    int    `json:"paymentID"`
	Bank         string `fjson:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Item struct {
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

type Subscriber struct {
	sub stan.Subscription
}

func New(clusterID, clientID string, subject string) *Subscriber {
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Println(err)
		return nil
	}
	sub, err := sc.Subscribe(subject, func(m *stan.Msg) {
		var order Order
		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			fmt.Println("[ERROR]", err)
		}
		fmt.Printf("Received a message\n")
		fmt.Println(order.OrderUID)
	})
	if err != nil {
		log.Println(err)
		return nil
	}
	return &Subscriber{sub: sub}
}

func (s *Subscriber) Unsubscribe() error {
	err := s.sub.Unsubscribe()
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

var (
	clusterID = "test-cluster"
	clientID  = "test"
)

type FakeJSONMessage struct {
	OrderUID          string `fake:"{regex:[a-z|0-9]{19}}" json:"order_uid"`
	TrackNumber       string `fake:"{regex:[A-Z]{13}}" json:"track_number"`
	Entry             string `fake:"{regex:[A-Z]{4}}" json:"entry"`
	Delivery          `json:"delivery"`
	Payment           `json:"payment"`
	Items             []*Item   `json:"items"`
	Locale            string    `fake:"{regex:[a-z]{2}}" json:"locale"`
	InternalSignature string    `fake:"skip" json:"internal_signature"`
	CustomerID        string    `fake:"{regex:[a-z|0-9]{19}}" json:"customer_id"`
	DeliveryService   string    `fake:"{regex:[a-z]{5}}" json:"delivery_service"`
	ShardKey          string    `fake:"{regex:[0-9]{1}}"  json:"shard_key"`
	SmID              int       `fake:"skip" json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `fake:"{regex:[0-9]{1}}" json:"oof_shard"`
}

type Delivery struct {
	Name    string `fake:"{name}" json:"name"`
	Phone   string `fake:"{phone}" json:"phone"`
	Zip     string `fake:"{regex:[0-9]{7}}" json:"zip"`
	City    string `fake:"{city}"`
	Address string `fake:"{street}" json:"address"`
	Region  string `fake:"{city}" json:"region"`
	Email   string `fake:"{email}" json:"email"`
}

type Payment struct {
	Transaction  string `fake:"{regex:[a-z|0-9]{19}}" json:"transaction"`
	RequestID    string `fake:"{regex:[a-z|0-9]{6}}" json:"request_id"`
	Currency     string `fake:"{regex:[a-z]{3}}" json:"currency"`
	Provider     string `fake:"{regex:[a-z]{5}}" json:"provider"`
	Amount       int    `fake:"skip" json:"amount"`
	PaymentID    int    `fake:"skip" json:"paymentID"`
	Bank         string `fake:"{regex:[a-z]{5}}" json:"bank"`
	DeliveryCost int    `fake:"skip" json:"delivery_cost"`
	GoodsTotal   int    `fake:"skip" json:"goods_total"`
	CustomFee    int    `fake:"skip" json:"custom_fee"`
}

type Item struct {
	ChrtID      int    `fake:"skip" json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `fake:"skip" json:"price"`
	Rid         string `fake:"{regex:[a-z|0-9]{19}}" json:"rid"`
	Name        string `fake:"{name}" json:"name"`
	Sale        int    `fake:"skip" json:"sale"`
	Size        string `fake:"{regex:[0-9]{2}}" json:"size"`
	TotalPrice  int    `fake:"skip" json:"total_price"`
	NmID        int    `fake:"skip" json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `fake:"skip" json:"status"`
}

const (
	subject = "orders"
)

func main() {
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatal(err)
	}
	for {
		var message FakeJSONMessage
		err := gofakeit.Struct(&message)
		if err != nil {
			log.Println("GOFAKEIT ERROR:", err)
		}
		message.Payment.Amount = gofakeit.Number(0, 100000000)
		message.Payment.PaymentID = gofakeit.Number(0, 100000000)
		message.Payment.DeliveryCost = gofakeit.Number(0, 100000)
		message.Payment.GoodsTotal = gofakeit.Number(0, 10000)
		message.Payment.CustomFee = gofakeit.Number(0, 10)
		message.SmID = gofakeit.Number(0, 200)
		for _, itm := range message.Items {
			itm.ChrtID = gofakeit.Number(0, 1000000)
			itm.Price = gofakeit.Number(0, 100000)
			itm.Sale = gofakeit.Number(0, 100)
			itm.TotalPrice = gofakeit.Number(0, 1000000)
			itm.NmID = gofakeit.Number(0, 1000000000)
			itm.Status = 202
		}

		data, err := json.Marshal(message)
		if err != nil {
			log.Fatal(err)
		}
		_ = sc.Publish(subject, data)
		log.Printf("send message in topic %s\n", subject)
		time.Sleep(5 * time.Second)
	}
}

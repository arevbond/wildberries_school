package cache

import (
	stan_sub "first_task_l0/internal/clients/stan-sub"
	"first_task_l0/internal/config"
	"first_task_l0/internal/storage"
	"first_task_l0/internal/storage/postgres"
	"github.com/pkg/errors"
)

const (
	CACHE_RECOVERY_LIMIT = 100
)

var (
	ErrOrderNotExists = errors.New("order doesn't exists")
)

type Cache struct {
	db storage.Storage
	mp map[string]stan_sub.Order
}

func New(cfg *config.Config) (*Cache, error) {
	db, err := postgres.New(cfg)
	if err != nil {
		return nil, err
	}
	mp := make(map[string]stan_sub.Order)
	return &Cache{mp: mp, db: db}, nil
}

func (c *Cache) Recovery() error {
	orders, err := c.db.GetOrders(CACHE_RECOVERY_LIMIT)
	if err != nil {
		return err
	}
	for _, order := range orders {
		customer, err := c.db.GetCustomer(order.CustomerID)
		if err != nil {
			return err
		}

		items, err := c.db.GetItems(order.OrderUID)
		if err != nil {
			return err
		}

		outDelivery := stan_sub.Delivery{
			Name:    customer.Name,
			Phone:   customer.Phone,
			Zip:     customer.Zip,
			City:    customer.City,
			Address: customer.Address,
			Region:  customer.Region,
			Email:   customer.Email,
		}

		outPayment := stan_sub.Payment{
			Transaction:  order.PaymentDB.Transaction,
			RequestID:    order.PaymentDB.RequestID,
			Currency:     order.PaymentDB.Currency,
			Provider:     order.PaymentDB.Provider,
			Amount:       order.PaymentDB.Amount,
			PaymentDT:    order.PaymentDB.PaymentDT,
			Bank:         order.PaymentDB.Bank,
			DeliveryCost: order.PaymentDB.DeliveryCost,
			GoodsTotal:   order.PaymentDB.GoodsTotal,
			CustomFee:    order.PaymentDB.CustomFee,
		}

		outItems := []stan_sub.Item{}
		for _, item := range items {
			outItem := stan_sub.Item{
				ChrtID:      item.ChrtID,
				TrackNumber: item.TrackNumber,
				Price:       item.Price,
				Rid:         item.Rid,
				Name:        item.Name,
				Sale:        item.Sale,
				Size:        item.Size,
				TotalPrice:  item.TotalPrice,
				NmID:        item.NmID,
				Brand:       item.Brand,
				Status:      item.Status,
			}
			outItems = append(outItems, outItem)
		}

		outOrder := stan_sub.Order{
			OrderUID:          order.OrderUID,
			TrackNumber:       order.TrackNumber,
			Entry:             order.Entry,
			Delivery:          outDelivery,
			Payment:           outPayment,
			Items:             outItems,
			Locale:            order.Locale,
			InternalSignature: order.InternalSignature,
			CustomerID:        order.CustomerID,
			DeliveryService:   order.DeliveryService,
			ShardKey:          order.ShardKey,
			SmID:              order.SmID,
			DateCreated:       order.DateCreated,
			OofShard:          order.OofShard,
		}

		c.CreateOrder(outOrder)
	}
	return nil
}

func (c *Cache) CreateOrder(order stan_sub.Order) {
	c.mp[order.OrderUID] = order
	return
}

func (c *Cache) GetOrder(uid string) (stan_sub.Order, error) {
	if order, ok := c.mp[uid]; ok {
		return order, nil
	}
	return stan_sub.Order{}, ErrOrderNotExists
}

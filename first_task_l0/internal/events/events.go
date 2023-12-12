package events

import (
	stan_sub "first_task_l0/internal/clients/stan-sub"
	"first_task_l0/internal/storage"
	"github.com/pkg/errors"
)

type Processor struct {
	db storage.Storage
}

func New(db storage.Storage) *Processor {
	return &Processor{
		db: db,
	}
}

// CreateOrder конвертирует приходящий JSON в типы таблиц базы данных
// и записывает в БД данные.
func (p *Processor) CreateOrder(inOrder stan_sub.Order) error {
	customer := &storage.CustomerDB{
		ID:      inOrder.CustomerID,
		Name:    inOrder.Delivery.Name,
		Phone:   inOrder.Delivery.Phone,
		Zip:     inOrder.Delivery.Zip,
		City:    inOrder.Delivery.City,
		Address: inOrder.Delivery.Address,
		Region:  inOrder.Delivery.Region,
		Email:   inOrder.Delivery.Email,
	}

	err := p.db.CreateCustomer(customer)
	if err != nil {
		return errors.Wrap(err, "can't create order in 'CreateOrder'")
	}

	payment := storage.PaymentDB{
		Transaction:  inOrder.Payment.Transaction,
		RequestID:    inOrder.Payment.RequestID,
		Currency:     inOrder.Payment.Currency,
		Provider:     inOrder.Payment.Provider,
		Amount:       inOrder.Payment.Amount,
		PaymentDT:    inOrder.Payment.PaymentID,
		Bank:         inOrder.Payment.Bank,
		DeliveryCost: inOrder.Payment.DeliveryCost,
		GoodsTotal:   inOrder.Payment.GoodsTotal,
		CustomFee:    inOrder.Payment.CustomFee,
	}

	order := &storage.OrderDB{
		ID:                inOrder.OrderUID,
		TrackNumber:       inOrder.TrackNumber,
		Entry:             inOrder.Entry,
		CustomerID:        inOrder.CustomerID,
		PaymentDB:         payment,
		Locale:            inOrder.Locale,
		InternalSignature: inOrder.InternalSignature,
		DeliveryService:   inOrder.DeliveryService,
		ShardKey:          inOrder.ShardKey,
		SmID:              inOrder.SmID,
		DateCreated:       inOrder.DateCreated,
		OofShard:          inOrder.OofShard,
	}
	err = p.db.CreateOrder(order)
	if err != nil {
		return errors.Wrap(err, "can't create order in 'CreateOrder':")
	}

	for _, itm := range inOrder.Items {
		item := &storage.ItemDB{
			ID:          itm.ChrtID,
			TrackNumber: inOrder.TrackNumber,
			Price:       itm.Price,
			Rid:         itm.Rid,
			Name:        itm.Name,
			Sale:        itm.Sale,
			TotalPrice:  itm.TotalPrice,
			NmID:        itm.NmID,
			Brand:       itm.Brand,
			Status:      itm.Status,
			OrderID:     inOrder.OrderUID,
		}
		if err = p.db.CreateItem(item); err != nil {
			return errors.Wrap(err, "can't create item in 'CreateOrder':")
		}
	}
	return nil
}

package order

import (
	stan_sub "first_task_l0/internal/clients/stan-sub"
	"first_task_l0/internal/storage/cache"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func New(db *cache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order stan_sub.Order
		var err error

		if orderID := chi.URLParam(r, "orderID"); orderID != "" {
			order, err = db.GetOrder(orderID)
			if err != nil {
				render.JSON(w, r, ErrNotFound)
			} else {
				render.JSON(w, r, order)
			}
		} else {
			render.Render(w, r, ErrNotFound)
			return
		}
	}
}

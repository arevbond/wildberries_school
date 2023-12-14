package display

import (
	"first_task_l0/internal/storage/cache"
	"html/template"
	"net/http"
)

func GetOrderById(db *cache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		ids := query["id"]
		var id string
		if len(ids) != 0 {
			id = ids[0]
		}
		order, _ := db.GetOrder(id)
		tmpl := template.Must(template.ParseFiles("internal/http_server/handlers/display/templates/index.html"))
		tmpl.Execute(w, order)
	}
}

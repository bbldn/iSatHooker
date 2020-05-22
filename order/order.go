package order

import "net/http"

func SynchronizeOrder(r *http.Request, path string) {
	if "Сохранить" == r.Form.Get("save") {
		//order:synchronize:by-ids r.Form.Get("id")
	}
}

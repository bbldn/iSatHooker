package category

import (
	"net/http"
)

func SynchronizeCategory(r *http.Request, path string) {
	if len(r.Form) > 1 {
		// category:synchronize:by-ids r.Form.Get("must_delete")
	}
}

func SynchronizeCategoryProducts(r *http.Request, path string) {
	if "Сохранить" == r.Form.Get("products_update") {
		//product:price:update:by-category-id r.Form.Get("categoryID")
		//product:synchronize:by-ids r.Form.Get("categoryID")
	}
}

package category

import (
	"fmt"
	"net/http"
	"net/url"
)

func SynchronizeCategory(r *http.Request, addr string) {
	if len(r.Form) > 1 {
		formData := url.Values{
			"command": {"category:synchronize:by-ids " + r.Form.Get("must_delete")},
		}

		_, err := http.PostForm(addr, formData)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func SynchronizeCategoryProducts(r *http.Request, addr string) {
	if "Сохранить" == r.Form.Get("products_update") {
		formData := url.Values{
			"command": {
				"product:price:update:by-category-id " + r.Form.Get("categoryID"),
				"product:synchronize:by-ids " + r.Form.Get("categoryID"),
			},
		}

		_, err := http.PostForm(addr, formData)
		if err != nil {
			fmt.Println(err)
		}

		formData = url.Values{
			"command": {"product:synchronize:by-ids " + r.Form.Get("categoryID")},
		}

		_, err = http.PostForm(addr, formData)
		if err != nil {
			fmt.Println(err)
		}
	}
}

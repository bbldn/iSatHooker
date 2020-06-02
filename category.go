package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c Context) SynchronizeCategory(r *http.Request) {
	if len(r.Form) > 1 {
		formData := url.Values{
			"command": {"category:synchronize:by-ids " + r.Form.Get("must_delete")},
		}

		_, err := http.PostForm(c.Config.Values["DEFERRED_OPERATIONS_ADDRESS"], formData)
		if err != nil {
			fmt.Println(err)
		}

		c.sendHook(r.Form, "HOOK_CATEGORY")
	}
}

func (c Context) SynchronizeCategoryProducts(r *http.Request) {
	if "Сохранить" == r.Form.Get("products_update") {
		formData := url.Values{
			"command": {
				"product:price:update:by-category-id " + r.Form.Get("categoryID"),
				"product:synchronize:by-ids " + r.Form.Get("categoryID"),
			},
		}

		_, err := http.PostForm(c.Config.Values["DEFERRED_OPERATIONS_ADDRESS"], formData)
		if err != nil {
			fmt.Println(err)
		}

		c.sendHook(r.Form, "HOOK_CATEGORY_PRODUCT")
	}
}

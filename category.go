package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c Context) SynchronizeCategory(r *http.Request) {
	if len(r.Form) > 1 {
		var formData map[string][]string
		if len(r.Form.Get("must_delete")) > 0 {
			formData = url.Values{
				"command": {fmt.Sprintf("category:synchronize:by-ids %s 1", r.Form.Get("must_delete"))},
			}
		} else if len(r.Form.Get("name")) > 0 {
			formData = url.Values{
				"command": {fmt.Sprintf("category:synchronize:by-name %s 1", r.Form.Get("name"))},
			}
		} else {
			return
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
				fmt.Sprintf("product:price:update:by-category-id %s", r.Form.Get("categoryID")),
				fmt.Sprintf("product:synchronize:by-ids %s 1", r.Form.Get("categoryID")),
			},
		}

		_, err := http.PostForm(c.Config.Values["DEFERRED_OPERATIONS_ADDRESS"], formData)
		if err != nil {
			fmt.Println(err)
		}

		c.sendHook(r.Form, "HOOK_CATEGORY_PRODUCT")
	}
}

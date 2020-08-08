package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func (c Context) SynchronizeCategory(r *http.Request) {
	if len(r.Form) > 1 {
		var formData map[string][]string
		var command string

		value := strings.TrimSpace(r.Form.Get("must_delete"))
		if 0 == len(value) {
			value = strings.TrimSpace(r.Form.Get("name"))
			if 0 == len(value) {
				return
			}

			command = "category:synchronize:by-name '" + value + "' 1"
		} else {
			command = fmt.Sprintf("category:synchronize:by-ids %s 1", value)
		}

		formData = url.Values{
			"command": {command},
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
		id := strings.TrimSpace(r.Form.Get("categoryID"))

		if 0 == len(id) {
			return
		}

		formData := url.Values{
			"command": {
				"currency:synchronize:all",
				fmt.Sprintf("product:synchronize:by-category-id %s 1", id),
			},
		}

		_, err := http.PostForm(c.Config.Values["DEFERRED_OPERATIONS_ADDRESS"], formData)
		if err != nil {
			fmt.Println(err)
		}

		c.sendHook(r.Form, "HOOK_CATEGORY_PRODUCT")
	}
}

package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c Context) SynchronizeGamePost(r *http.Request) {
	switch r.Form.Get("sub") {
	case "main":
		c.synchronizeMain(r)
	case "price":
		c.synchronizePrice(r)
	case "orders":
	default:
		c.synchronizeOrder(r)
	}
}

func (c Context) synchronizeOrder(r *http.Request) {
	if "Сохранить" == r.Form.Get("save") {
		formData := url.Values{
			"command": {fmt.Sprintf("order:synchronize:by-ids backToFront %s", r.Form.Get("id"))},
		}

		_, err := http.PostForm(c.Config.Values["DEFERRED_OPERATIONS_ADDRESS"], formData)
		if err != nil {
			fmt.Println(err)
		}

		c.sendHook(r.Form, "HOOK_ORDER")
	}
}

func (c Context) synchronizePrice(r *http.Request) {
	formData := url.Values{
		"command": {"product:price:update:all"},
	}

	_, err := http.PostForm(c.Config.Values["DEFERRED_OPERATIONS_ADDRESS"], formData)
	if err != nil {
		fmt.Println(err)
	}

	c.sendHook(r.Form, "HOOK_ORDER")
}

func (c Context) synchronizeMain(r *http.Request) {
	if "currencies" == r.Form.Get("setting") && "Сохранить" == r.Form.Get("currency_save") {
		formData := url.Values{
			"command": {
				"currency:synchronize:all",
				"product:price:update:all",
			},
		}

		_, err := http.PostForm(c.Config.Values["DEFERRED_OPERATIONS_ADDRESS"], formData)
		if err != nil {
			fmt.Println(err)
		}

		c.sendHook(r.Form, "HOOK_CURRENCY")
	}
}

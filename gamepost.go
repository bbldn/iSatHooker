package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func (c Context) SynchronizeGamePost(r *http.Request) {
	switch r.Form.Get("sub") {
	case "main":
		c.synchronizeMain(r)
	case "price":
		c.synchronizePrice(r)
	case "orders":
		c.synchronizeOrder(r)
	case "club":
		c.synchronizeClub(r)
	}
}

func (c Context) synchronizeOrder(r *http.Request) {
	if "Сохранить" == r.Form.Get("save") {
		id := strings.TrimSpace(r.Form.Get("id"))
		if 0 == len(id) {
			return
		}

		formData := url.Values{
			"command": {fmt.Sprintf("order:synchronize:by-ids backToFront %s", id)},
		}

		_, err := http.PostForm(c.Config.Values["DEFERRED_OPERATIONS_ADDRESS"], formData)
		if err != nil {
			fmt.Println(err)
		}

		c.sendHook(r.Form, "HOOK_ORDER")
	}
}

func (c Context) synchronizePrice(r *http.Request) {
	if "Сохранить" == r.Form.Get("prices_save") {
		formData := url.Values{
			"command": {
				"currency:synchronize:all",
				"product:price:synchronize:all:fast",
			},
		}

		_, err := http.PostForm(c.Config.Values["DEFERRED_OPERATIONS_ADDRESS"], formData)
		if err != nil {
			fmt.Println(err)
		}

		c.sendHook(r.Form, "HOOK_PRICE_ALL")
	}
}

func (c Context) synchronizeMain(r *http.Request) {
	if "currencies" == r.Form.Get("setting") && "Сохранить" == r.Form.Get("currency_save") {
		formData := url.Values{
			"command": {
				"currency:synchronize:all",
				"product:price:synchronize:all:fast",
			},
		}

		_, err := http.PostForm(c.Config.Values["DEFERRED_OPERATIONS_ADDRESS"], formData)
		if err != nil {
			fmt.Println(err)
		}

		c.sendHook(r.Form, "HOOK_CURRENCY")
	}
}

func (c Context) synchronizeClub(r *http.Request) {
	if "Обновить" == r.Form.Get("save") {
		id := strings.TrimSpace(r.Form.Get("id"))
		if 0 == len(id) {
			return
		}

		formData := url.Values{
			"command": {fmt.Sprintf("customer:synchronize:by-ids %s", id)},
		}

		_, err := http.PostForm(c.Config.Values["DEFERRED_OPERATIONS_ADDRESS"], formData)
		if err != nil {
			fmt.Println(err)
		}

		c.sendHook(r.Form, "HOOK_CLUB")
	}
}

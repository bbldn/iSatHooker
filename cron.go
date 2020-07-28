package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c Context) SynchronizePrices(r *http.Request, addr string) {
	formData := url.Values{
		"command": {
			"currency:synchronize:all",
			"product:price:synchronize:all",
		},
	}

	_, err := http.PostForm(c.Config.Values["DEFERRED_OPERATIONS_ADDRESS"], formData)
	if err != nil {
		fmt.Println(err)
	}

	c.sendHook(r.Form, "HOOK_CRON_PRICES")
}

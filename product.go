package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c Context) SynchronizeProduct(r *http.Request, addr string) {
	if len(r.Form) > 1 {
		formData := url.Values{
			"command": {"product:synchronize:by-ids " + r.Form.Get("productID")},
		}

		_, err := http.PostForm(c.Config.Values["DEFERRED_OPERATIONS_ADDRESS"], formData)
		if err != nil {
			fmt.Println(err)
		}

		c.sendHook(r.Form, "HOOK_PRODUCT")
	}
}

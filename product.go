package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func (c Context) SynchronizeProduct(r *http.Request, addr string) {
	if len(r.Form) > 1 {
		var command string

		value := strings.TrimSpace(r.Form.Get("productID"))
		if 0 == len(value) || "0" == value {
			command = "product:synchronize:last 1"
		} else {
			command = fmt.Sprintf("product:synchronize:by-ids %s 1", value)
		}

		formData := url.Values{
			"command": {command},
		}

		_, err := http.PostForm(c.Config.Values["DEFERRED_OPERATIONS_ADDRESS"], formData)
		if err != nil {
			fmt.Println(err)
		}

		c.sendHook(r.Form, "HOOK_PRODUCT")
	}
}

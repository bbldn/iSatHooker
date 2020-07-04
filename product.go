package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c Context) SynchronizeProduct(r *http.Request, addr string) {
	if len(r.Form) > 1 {
		var formData map[string][]string

		if len(r.Form.Get("productID")) > 0 && "0" != r.Form.Get("productID") {
			formData = url.Values{
				"command": {fmt.Sprintf("product:synchronize:by-ids %s 1", r.Form.Get("productID"))},
			}
		} else if len(r.Form.Get("name")) > 0 {
			formData = url.Values{
				"command": {fmt.Sprintf("product:synchronize:by-name '%s' 1", r.Form.Get("name"))},
			}
		} else {
			return
		}

		_, err := http.PostForm(c.Config.Values["DEFERRED_OPERATIONS_ADDRESS"], formData)
		if err != nil {
			fmt.Println(err)
		}

		c.sendHook(r.Form, "HOOK_PRODUCT")
	}
}

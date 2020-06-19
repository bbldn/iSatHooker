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
				"command": {"product:synchronize:by-ids " + r.Form.Get("productID") + " 1"},
			}
		} else if len(r.Form.Get("name")) > 0 {
			formData = url.Values{
				"command": {"product:synchronize:by-name " + r.Form.Get("name") + " 1"},
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

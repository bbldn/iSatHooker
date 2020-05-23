package product

import (
	"fmt"
	"net/http"
	"net/url"
)

func SynchronizeProduct(r *http.Request, addr string) {
	if len(r.Form) > 1 {
		formData := url.Values{
			"command": {"product:synchronize:by-ids " + r.Form.Get("must_delete")},
		}

		_, err := http.PostForm(addr, formData)
		if err != nil {
			fmt.Println(err)
		}
	}
}

package order

import (
	"fmt"
	"net/http"
	"net/url"
)

func SynchronizeOrder(r *http.Request, addr string) {
	if "Сохранить" == r.Form.Get("save") {
		formData := url.Values{
			"command ": {"order:synchronize:by-ids " + r.Form.Get("id")},
		}

		_, err := http.PostForm(addr, formData)
		if err != nil {
			fmt.Println(err)
		}
	}
}

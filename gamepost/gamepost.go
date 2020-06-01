package gamepost

import (
	"fmt"
	"net/http"
	"net/url"
)

func Synchronize(r *http.Request, addr string) {
	switch r.Form.Get("sub") {
	case "orders":
		synchronizeOrder(r, addr)
	case "main":
		synchronizeMain(r, addr)
	}
}

func synchronizeOrder(r *http.Request, addr string) {
	if "Сохранить" == r.Form.Get("save") {
		formData := url.Values{
			"command": {"order:synchronize:by-ids backToFront " + r.Form.Get("id")},
		}

		_, err := http.PostForm(addr, formData)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func synchronizeMain(r *http.Request, addr string) {
	if "currencies" == r.Form.Get("setting") && "Сохранить" == r.Form.Get("currency_save") {
		formData := url.Values{
			"command": {"currency:synchronize:all"},
		}

		_, err := http.PostForm(addr, formData)
		if err != nil {
			fmt.Println(err)
		}
	}
}

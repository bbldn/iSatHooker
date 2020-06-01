package main

import (
	"fmt"
	"isatHooker/category"
	"isatHooker/config"
	"isatHooker/gamepost"
	"isatHooker/product"
	"isatHooker/response"
	"net/http"
)

var appConfig config.Config

func start(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		_, _ = fmt.Fprintf(w, response.Response{Ok: true, Errors: []string{"Error parse body"}}.ToJson())

		return false
	}

	_, _ = fmt.Fprintf(w, response.Response{Ok: true}.ToJson())

	return true
}

func categoryUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if false == start(w, r) {
		return
	}

	go category.SynchronizeCategory(r, appConfig.Values["DEFERRED_OPERATIONS_ADDRESS"])
}

func catalogUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if false == start(w, r) {
		return
	}

	dpt := r.Form.Get("dpt")

	switch dpt {
	case "catalog":
		go category.SynchronizeCategoryProducts(r, appConfig.Values["DEFERRED_OPERATIONS_ADDRESS"])
	case "gamepost":
		go gamepost.Synchronize(r, appConfig.Values["DEFERRED_OPERATIONS_ADDRESS"])
	}
}

func productUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if false == start(w, r) {
		return
	}

	go product.SynchronizeProduct(r, appConfig.Values["DEFERRED_OPERATIONS_ADDRESS"])
}

func main() {
	err := appConfig.Load(make(map[string]string))
	if nil != err {
		fmt.Println(err)

		return
	}

	http.HandleFunc("/back/product/update", productUpdateHandler)
	http.HandleFunc("/back/category/update", categoryUpdateHandler)
	http.HandleFunc("/back/admin/update", catalogUpdateHandler)

	addr := fmt.Sprintf("%s:%s", appConfig.Values["ADDRESS"], appConfig.Values["PORT"])
	err = http.ListenAndServe(addr, nil)
	if nil != err {
		fmt.Println(err)
	}
}

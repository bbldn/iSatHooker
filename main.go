package main

import (
	"fmt"
	"isatHooker/config"
	"isatHooker/response"
	"net/http"
)

type Context struct {
	Config config.Config
}

var context Context

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

	go context.SynchronizeCategory(r)
}

func catalogUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if false == start(w, r) {
		return
	}

	dpt := r.Form.Get("dpt")

	switch dpt {
	case "catalog":
		go context.SynchronizeCategoryProducts(r)
	case "gamepost":
		go context.SynchronizeGamePost(r)
	}
}

func productUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if false == start(w, r) {
		return
	}

	go context.SynchronizeProduct(r, context.Config.Values["DEFERRED_OPERATIONS_ADDRESS"])
}

func main() {
	context = Context{Config: config.Config{}}
	err := context.Config.Load(make(map[string]string))
	if nil != err {
		fmt.Println(err)

		return
	}

	http.HandleFunc("/back/product/update", productUpdateHandler)
	http.HandleFunc("/back/category/update", categoryUpdateHandler)
	http.HandleFunc("/back/admin/update", catalogUpdateHandler)

	addr := fmt.Sprintf("%s:%s", context.Config.Values["ADDRESS"], context.Config.Values["PORT"])
	err = http.ListenAndServe(addr, nil)
	if nil != err {
		fmt.Println(err)
	}
}

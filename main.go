package main

import (
	"fmt"
	"isatHooker/category"
	"isatHooker/response"
	"net/http"
)

var path string

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

	go category.SynchronizeCategory(r, path)
}

func main() {
	path = "/home/user/PhpstormProjects/iSatSynchronizer/bin/console"

	http.HandleFunc("/back/category/update", categoryUpdateHandler)

	addr := fmt.Sprintf("%s:%s", "0.0.0.0", "8082")
	_ = http.ListenAndServe(addr, nil)
}

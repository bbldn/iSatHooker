package category

import (
	"fmt"
	"net/http"
	"os/exec"
)

func SynchronizeCategory(r *http.Request, path string) {
	if len(r.Form) > 1 {
		commands := []string{path, "category:synchronize:by-ids", r.Form.Get("must_delete")}
		cmd := exec.Command("php", commands...)
		_ = cmd.Start()
	}
}

func SynchronizeCategoryProducts(r *http.Request, path string) {
	if r.Form.Get("products_update") == "Сохранить" {
		fmt.Println(r.Form.Get("categoryID"))
		//commands := []string{path, "category:synchronize:by-ids", r.Form.Get("must_delete")}
		//cmd := exec.Command("php", commands...)
		//_ = cmd.Start()
	}
}

package main

import (
	"fmt"
	"net/http"
	"strings"
)

func (c Context) sendHook(values map[string][]string, hook string) {
	addresses, exists := c.Config.Values[hook]
	if false == exists {
		return
	}

	parsedAddresses := strings.Split(addresses, ",")
	for _, address := range parsedAddresses {
		_, err := http.PostForm(address, values)
		if err != nil {
			fmt.Println(err)
		}
	}
}

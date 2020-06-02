package main

import (
	"fmt"
	"net/http"
	"strings"
)

func (c Context) sendHook(values map[string][]string, hook string) {
	addresses := c.Config.Values[hook]
	parsedAddresses := strings.Split(addresses, ",")
	for _, address := range parsedAddresses {
		_, err := http.PostForm(address, values)
		if err != nil {
			fmt.Println(err)
		}
	}
}

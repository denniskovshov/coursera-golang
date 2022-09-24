package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var addr string

	fmt.Println("Enter a name:")
	fmt.Scan(&name)

	fmt.Println("Enter an address:")
	fmt.Scan(&addr)

	data := make(map[string]string)

	data["name"] = name
	data["address"] = addr

	jsonData, _ := json.Marshal(data)

	fmt.Println(string(jsonData))
}

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	

	x := make(map[string]string)
	var name, add string
	fmt.Println("enter the name")
	fmt.Scan(&name)
	x["name"] = name

	fmt.Println("enter the address")
	fmt.Scan(&add)
	x["address"] = add

	js, _ := json.MarshalIndent(x, " ", "  ")
	fmt.Println("-----printing the json object----")
	print(string(js))
}

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var inputName string
	var inputAddress string

	fmt.Println("Enter the Name")
	_, err := fmt.Scan(&inputName)
	if err != nil {
		return
	}

	fmt.Println("Enter the Address")
	_, err2 := fmt.Scan(&inputAddress)
	if err2 != nil {
		return
	}

	pairMap := make(map[string]string, 1)
	pairMap[inputName] = inputAddress

	marshal, err3 := json.Marshal(pairMap)
	if err3 != nil {
		return
	}
	fmt.Println(string(marshal[:]))
}

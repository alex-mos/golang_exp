package main

import "fmt"

func main() {
	//numbers := []int{0, 1, 2, 3, 4, 5}
	//fmt.Println(numbers[1:4])

	var user map[string]bool = map[string]bool{
		"firstName": true,
		"lastName":  false,
	}

	if _, keyExist := user["lastName"]; keyExist {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	if _, keyExist := user["pastName"]; keyExist {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	//fmt.Println(user["firstName"])
	//fmt.Println(user["lastName"])
	//fmt.Println(user["pastName"])
}

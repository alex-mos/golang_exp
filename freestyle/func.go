package main

import "fmt"

func callMyName(name string) (string, error) {
	return name, nil
}

func main() {
	fmt.Println(callMyName("User"))
}
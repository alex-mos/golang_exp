package main

import "fmt"

func main() {
	s := []string{"alex", "anya", "jin"}

	for index, value := range s {
		fmt.Println("%v: %v", index, value)
	}
}
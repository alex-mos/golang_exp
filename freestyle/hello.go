package main

import "fmt"

func main() {
	var a1 [3]int
	fmt.Println("a1 short", a1)
	fmt.Println("a1 short %v\n", a1)
	fmt.Println("a1 short %#v\n", a1)

	const size = 2

	var a2 [2 * size]bool // [false,false,false,false]
	fmt.Println("a2", a2)
	fmt.Println(a2)

	a3 := [...]int{1,5,3}
	fmt.Println(a3)

	// вместо массива, можно использовать структуру данных слайс:

	var buf0 []int
	buf1 := []int{}

	println(buf0)
	println(buf1)

}
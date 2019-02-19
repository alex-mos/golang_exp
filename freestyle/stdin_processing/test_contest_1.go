package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	//a := 0

	for input.Scan() {
		fmt.Println(input.Text())
		//newInt, err := strconv.Atoi(input.Text())
		//if err != nil {
		//	panic(err)
		//}
		//a += newInt
	}
	fmt.Println("end of file")
	//fmt.Println(a)
}

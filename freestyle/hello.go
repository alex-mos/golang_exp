package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)

	fmt.Println(stdin.Text())

	//for stdin.Scan() {
	//	fmt.Println(stdin.Text())
	//}
}

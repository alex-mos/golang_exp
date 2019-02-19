package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	b := make([]byte, 0)
	b = append(b, 5)
	b = append(b, 120)
	b = append(b, 255)
	b = append(b, 255)

	fmt.Println(b)

	bb := bytes.NewBuffer(b)

	fmt.Fprintln(os.Stdout, bb)
}

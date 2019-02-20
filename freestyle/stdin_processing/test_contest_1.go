package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	str, err := ioutil.ReadAll(os.Stdin)
	fmt.Fprintf(os.Stdout, "string: %s", str)
	fmt.Fprintf(os.Stdout, "length: %d\n", len(str))
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %s\n", err)
	}
}

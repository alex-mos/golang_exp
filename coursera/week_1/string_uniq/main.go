package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func uniq(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input) // создаём новый сканнер из стандартного ввода
	var prev string
	for in.Scan() {
		txt := in.Text()
		if txt == prev {
			continue
		}
		if txt < prev {
			return fmt.Errorf("file is not sorted")
		}
		prev = txt
		fmt.Fprintln(output, txt) // Принт, который принимает на вход io.Reader
	}
	return nil
}

func main() {
	err := uniq(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
)

func dirTree(out io.Writer, path string, printFiles bool) error {
	items, err := ioutil.ReadDir(path)
	if err != nil {
		return fmt.Errorf("Path reading error")
	}

	if printFiles == false {
		items = filterDirs(items)
	}

	for index, item := range items {
		if item.IsDir() {
			if index < len(items) - 1 {
				fmt.Printf("├───")
			} else {
				fmt.Printf("└───")
			}
			fmt.Printf("%v\n", item.Name())
		}

	}

	return nil
}

func filterDirs(items []os.FileInfo) []os.FileInfo {
	var result []os.FileInfo
	for _, item := range items {
		if item.IsDir() {
			result = append(result, item)
		}
	}
	return result
}

// todo: разобраться, почему не работает, и что сюда передаётся – ссылки или значения
func itemSort(items []os.FileInfo) []os.FileInfo {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Name() < items[j].Name()
	})
	return items
}

func main() {
	var out = os.Stdout

	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}

	var path string = os.Args[1]
	var printFiles bool = len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, path, printFiles)

	if err != nil {
		panic(err.Error())
	}
}

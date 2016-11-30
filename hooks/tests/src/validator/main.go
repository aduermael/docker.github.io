package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"unicode"
)

// type of function that perform validation on a single file
type FileTestFunc func(path string) error

var (
	test []FileTestFunc = []FileTestFunc{
		markdownHasMainTitle,
		frontmatterHasTitle,
	}
)

func main() {
	fmt.Println("-----------------")
	fmt.Println("RUN TESTS")
	fmt.Println("-----------------")

	//
	for i, fn := range test {
		fmt.Printf("--- test %d\n", i+1)
		filepath.Walk("/docs", func(path string, info os.FileInfo, err error) error {
			err = fn(path)
			if err != nil {
				fmt.Println(err.Error(), "-", path)
				// os.Exit(1)
			}
			return nil
		})
	}

	fmt.Println("----------------------")
	fmt.Println("ALL CHECKS HAVE PASSED")
	fmt.Println("----------------------")
}

// --------------------
// FILE EXTENSIONS
// --------------------

func lowercaseExtension(path string) error {
	ext := filepath.Ext(path)
	if len(ext) > 0 {
		for _, r := range ext {
			if unicode.IsUpper(r) {
				fmt.Println("lowercaseExtension", path)
				return errors.New(":(")
			}
		}
	}
	return nil
}

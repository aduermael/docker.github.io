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
		lowercaseExtension,
	}
)

func main() {
	fmt.Println("-----------------")
	fmt.Println("RUN TESTS")
	fmt.Println("-----------------")

	var err error

	//
	for i, fn := range test {
		fmt.Printf("--- test %d\n", i+1)
		err = filepath.Walk("/docs", func(path string, info os.FileInfo, err error) error {
			return fn(path)
		})
		if err != nil {
			fmt.Println("----------------------")
			fmt.Println("TEST FAILED")
			fmt.Println(err.Error())
			fmt.Println("----------------------")
			os.Exit(1)
		}
	}

	fmt.Println("----------------------")
	fmt.Println("ALL CHECKS HAVE PASSED")
	fmt.Println("----------------------")
}

//
//
// FILE EXTENSIONS
//
//

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

package main

import (
	"bytes"
	"fmt"
	"github.com/gdevillele/frontparser"
	"io/ioutil"
	"strings"
)

func markdownHasMainTitle(path string) error {
	// if file has .md extension
	if strings.HasSuffix(path, ".md") {
		fileBytes, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		// if file has frontmatter
		if frontparser.HasFrontmatterHeader(fileBytes) {
			fm, md, err := frontparser.ParseFrontmatterAndContent(fileBytes)
			if err != nil {
				return err
			}
			// if frontmatter contains a "published" value
			if _ /*title*/, exists := fm["title"]; exists {
				md = bytes.TrimSpace(md)
				if bytes.HasPrefix(md, []byte("# ")) {
					fmt.Println("markdown has top-level title", path)
				}
			}
		} else { // no frontmatter
			return nil
		}
	}
	return nil
}

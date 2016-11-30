package main

import (
	"bytes"
	"errors"
	"github.com/gdevillele/frontparser"
	"io/ioutil"
	"strings"
)

func frontmatterHasTitle(path string) error {
	if strings.HasSuffix(path, ".md") {
		fileBytes, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		// if file has frontmatter
		if frontparser.HasFrontmatterHeader(fileBytes) {
			fm, _, err := frontparser.ParseFrontmatterAndContent(fileBytes)
			if err != nil {
				return err
			}

			// skip markdowns that are not published
			if published, exists := fm["published"]; exists {
				if publishedBool, ok := published.(bool); ok {
					if publishedBool == false {
						return nil
					}
				}
			}

			if _, exists := fm["title"]; exists == false {
				return errors.New("can't find title in frontmatter")
			}
		} else { // no frontmatter
			return errors.New("can't find frontmatter")
		}
	}
	return nil
}

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

			if published, exists := fm["published"]; exists {
				if publishedBool, ok := published.(bool); ok {
					if publishedBool == false {
						// page is not published, we can skip
						return nil
					}
				}
			}

			if _ /*title*/, exists := fm["title"]; exists {
				md = bytes.TrimSpace(md)
				if bytes.HasPrefix(md, []byte("# ")) {
					return errors.New("markdown has top-level title")
				}
			}
		} else { // no frontmatter
			return nil
		}
	}
	return nil
}

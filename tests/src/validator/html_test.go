package main

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

var countLinks = 0
var countImages = 0

// TestURLs tests if we're not using absolute paths for URLs
// when pointing to local pages.
func TestURLs(t *testing.T) {
	count := 0

	filepath.Walk("/usr/src/app/allvbuild", func(path string, info os.FileInfo, err error) error {

		relPath := strings.TrimPrefix(path, "/usr/src/app/allvbuild")

		isArchive, err := regexp.MatchString(`^/v[0-9]+\.[0-9]+/.*`, relPath)
		if err != nil {
			t.Error(err.Error(), "-", relPath)
		}
		// skip archives for now, only test URLs in current version
		// TODO: test archives
		if isArchive {
			return nil
		}

		if err != nil {
			t.Error(err.Error(), "-", relPath)
		}
		b, htmlBytes, err := isHTML(path)
		if err != nil {
			t.Error(err.Error(), "-", relPath)
		}
		// don't check non-html files
		if b == false {
			return nil
		}

		count++

		err = testURLs(htmlBytes, path)
		if err != nil {
			t.Error(err.Error(), "-", relPath)
		}
		return nil
	})

	fmt.Println("found", count, "html files (excluding archives)")
	fmt.Println("found", countLinks, "links (excluding archives)")
	fmt.Println("found", countImages, "images (excluding archives)")
}

// testURLs tests if we're not using absolute paths for URLs
// when pointing to local pages.
func testURLs(htmlBytes []byte, htmlPath string) error {

	reader := bytes.NewReader(htmlBytes)

	z := html.NewTokenizer(reader)

	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			// End of the document, we're done
			return nil
		case html.StartTagToken:
			t := z.Token()

			urlStr := ""

			// check tag types
			switch t.Data {
			case "a":
				countLinks++
				ok, href := getHref(t)
				// skip, it may just be an anchor
				if !ok {
					break
				}
				urlStr = href

			case "img":
				countImages++
				ok, src := getSrc(t)
				if !ok {
					return errors.New("img with no src: " + t.String())
				}
				urlStr = src
			}

			// there's an url to test!
			if urlStr != "" {
				u, err := url.Parse(urlStr)
				if err != nil {
					return errors.New("can't parse url: " + t.String())
				}
				// test with github.com
				if u.Scheme != "" && u.Host == "docs.docker.com" {
					return errors.New("found absolute link: " + t.String())
				}

				// relative link
				if u.Scheme == "" {
					p := filepath.Join(htmlPath, mdToHtmlPath(u.Path))
					if _, err := os.Stat(p); os.IsNotExist(err) {

						fail := true

						// index.html could mean there's a corresponding index.md meaning built the correct path
						// but Jekyll actually creates index.html files for all md files.
						// foo.md -> foo/index.html
						// it does this to prettify urls, content of foo.md would then be rendered here:
						// http://domain.com/foo/ (instead of http://domain.com/foo.html)
						// so if there's an error, let's see if index.md exists, otherwise retry from parent folder
						if filepath.Base(htmlPath) == "index.html" {
							// retry from parent folder
							p = filepath.Join(filepath.Dir(htmlPath), "..", mdToHtmlPath(u.Path))
							if _, err := os.Stat(p); err == nil {
								fail = false
							}
						}

						if fail {
							return errors.New("relative link to non-existent resource: " + t.String())
						}
					}
				}
			}
		}
	}

	return nil
}

func mdToHtmlPath(mdPath string) string {
	if strings.HasSuffix(mdPath, ".md") == false {
		// file is not a markdown, don't change anything
		return mdPath
	}
	if strings.HasSuffix(mdPath, "index.md") {
		return strings.TrimSuffix(mdPath, "md") + "html"
	}
	return strings.TrimSuffix(mdPath, ".md") + "/index.html"
}

// helpers

func getHref(t html.Token) (ok bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}
	return
}

func getSrc(t html.Token) (ok bool, src string) {
	for _, a := range t.Attr {
		if a.Key == "src" {
			src = a.Val
			ok = true
		}
	}
	return
}

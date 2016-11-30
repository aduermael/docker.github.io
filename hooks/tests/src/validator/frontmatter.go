package main

import (
// "github.com/gdevillele/frontparser"
// "io/ioutil"
)

// func frontmatterPresentInMarkdownFiles(path string) error {

// }

// import (
// 	"bytes"
// 	"errors"
// 	"fmt"
// 	"gopkg.in/yaml.v2"
// 	"io/ioutil"
// 	"strings"

// 	"github.com/Sirupsen/logrus"
// 	// "github.com/gdevillele/frontparser"
// 	// "path/filepath"
// 	// "strings"
// )

// // validates that frontmatter, if present, is in YAML format
// func validateFrontmatterYamlFormat(path string) error {
// 	fileBytes, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		return err
// 	}
// 	fileBytes = bytes.TrimSpace(fileBytes)

// 	// test if YAML frontmatter
// 	if bytes.HasPrefix(fileBytes, []byte("---")) {

// 	} else if bytes.HasPrefix(fileBytes, []byte("+++")) {
// 		return errors.New("frontmatter looks to be TOML")
// 	} else if bytes.HasPrefix(fileBytes, []byte("{")) {
// 		return errors.New("frontmatter looks to be JSON")
// 	}
// 	return nil
// }

// func searchAndRepairKeywords(path string) error {

// 	// only consider paths ending with .md
// 	if strings.HasSuffix(path, ".md") == false {
// 		return nil
// 	}

// 	// check if path points to a regular file
// 	exists := regularFileExists(path)
// 	if exists == false {
// 		return nil
// 	}

// 	// read file's content
// 	fileBytes, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		return err
// 	}

// 	//
// 	if frontparser.HasFrontmatterHeader(fileBytes) == false {
// 		return nil
// 	}

// 	//
// 	fmMap, rest, err := frontparser.ParseFrontmatterAndContent(fileBytes)
// 	if err != nil {
// 		return err
// 	}

// 	//
// 	if _, exists := fmMap["keywords"]; exists == false {
// 		return nil
// 	}

// 	//
// 	keywordArray, err := interfaceConv.ToStringArray(fmMap["keywords"])
// 	if err != nil {
// 		logrus.Println("ERROR 1 (keywords is not an array)", err.Error())
// 		return nil
// 	}

// 	// keywords has a length different than 1
// 	if len(keywordArray) != 1 {
// 		logrus.Println("ERROR 2 (keywords has a length different than 1)")
// 		return nil
// 	}

// 	// re-generate md file
// 	fmMap["keywords"] = keywordArray[0]
// 	yamlBytes, err := yaml.Marshal(fmMap)
// 	if err != nil {
// 		logrus.Println("ERROR 3 (cannot yaml marshal)", err.Error())
// 		return err
// 	}

// 	newFileBytes := fmt.Sprintf("---\n")
// 	newFileBytes += string(yamlBytes)
// 	newFileBytes += fmt.Sprintf("---")
// 	newFileBytes += string(rest)

// 	err = ioutil.WriteFile(path, []byte(newFileBytes), os.ModePerm)
// 	if err != nil {
// 		logrus.Println("ERROR 4 (cannot write file)", err.Error())
// 		return err
// 	}

// 	return nil
// }

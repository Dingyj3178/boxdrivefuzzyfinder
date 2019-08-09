package main

import (
	"io/ioutil"
	"strings"

	"github.com/sahilm/fuzzy"
)

var filenamesBytes []byte
var err error
var filenames []string

func FileReader(path string) []string {
	filenamesBytes, err = ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	filenames = strings.Split(string(filenamesBytes), "\n")
	return filenames
}

func finder(path, pattern string) []string {
	f := FileReader(path)
	var fileList []string
	matches := fuzzy.Find(pattern, f)
	for _, match := range matches {
		fileList = append(fileList, string(match.Str))
		// fmt.Print(string(match.Str))
		// for i := 0; i < len(match.Str); i++ {
		// 	fmt.Print(string(match.Str[i]))
		// }
		// fmt.Println()
	}
	return fileList
}

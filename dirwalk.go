package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
)

func dirwalk(dir string) {
	iFile := getIgnore()
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			println(file.Name(), "is dir, Scan dir. ")
			if checkignore(file.Name(), iFile) {
				dirwalk(filepath.Join(dir, file.Name()))
				continue
			} else {
				println(file.Name(), "is ignored")
			}
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
		err = writeBuffering(paths)
		if err != nil {
			panic(err)
		}
	}
	// return paths
}

func getIgnore() []string {
	iFile := fileReader("./.boxignore")
	return iFile
}

func checkignore(file string, iFile []string) bool {
	for _, v := range iFile {
		if file == v {
			return false
		}
	}
	return true
}

func writeBuffering(filenames []string) error {
	file, err := os.OpenFile("./files.in", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, v := range filenames {
		if _, err := writer.WriteString(v + "\n"); err != nil {
			panic(err)
		}
	}
	err = writer.Flush()
	if err != nil {
		panic(err)
	}

	return nil
}

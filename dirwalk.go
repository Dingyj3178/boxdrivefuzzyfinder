package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
)

func dirwalk(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			println(file.Name(), "is dir, Scan dir. ")
			dirwalk(filepath.Join(dir, file.Name()))
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
		writeBuffering(paths)
	}

	// return paths
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
	writer.Flush()

	return nil
}

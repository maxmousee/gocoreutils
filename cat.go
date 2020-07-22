package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//ArgIsFile checks if a string is potentially a file path
func ArgIsFile(arg string) bool {
	return !strings.HasPrefix(arg, "-")
}

//FileExists checks if the file exists and user has permission to read it
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//GetFilesToRead Returns a list of files from os.Args
func GetFilesToRead() []string {
	argsWithoutProg := os.Args[1:]
	inputFiles := []string{}
	for _, currentFile := range argsWithoutProg {
		if ArgIsFile(currentFile) && FileExists(currentFile) {
			inputFiles = append(inputFiles, currentFile)
		}
	}
	return inputFiles
}

//ReadFile returns a string with file contents
func ReadFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(content)
}

func main() {
	inputFiles := GetFilesToRead()
	for _, currentFile := range inputFiles {
		fmt.Println(ReadFile(currentFile))
	}
}

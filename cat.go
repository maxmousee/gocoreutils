package main

import (
	"fmt"
	"os"
)

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

func main() {
	inputFiles := GetFilesToRead()
	for _, currentFile := range inputFiles {
		fmt.Println(ReadFile(currentFile))
	}
}

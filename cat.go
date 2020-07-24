package main

import (
	"fmt"
	"os"
)

func getFilesToRead() []string {
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
	inputFiles := getFilesToRead()
	for _, currentFile := range inputFiles {
		fmt.Println(ReadFile(currentFile))
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func isFile(arg string) bool {
	return !strings.HasPrefix(arg, "-")
}

func getFiles() []string {
	argsWithoutProg := os.Args[1:]
	inputFiles := make([]string, len(argsWithoutProg))
	for i, s := range argsWithoutProg {
		if isFile(s) {
			inputFiles[i] = s
		}
	}
	return inputFiles
}

func main() {
	inputFiles := getFiles()
	fmt.Println(inputFiles)
}

package main

import (
	"fmt"
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
	inputFiles := make([]string, len(argsWithoutProg))
	for i, s := range argsWithoutProg {
		if ArgIsFile(s) && FileExists(s) {
			inputFiles[i] = s
		}
	}
	return inputFiles
}

// func

func main() {
	inputFiles := GetFilesToRead()
	fmt.Println(inputFiles)
}

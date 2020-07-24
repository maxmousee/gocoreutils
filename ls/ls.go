package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//ArgIsDir checks if a string is potentially a directory path
func ArgIsDir(arg string) bool {
	return !strings.HasPrefix(arg, "-")
}

//DirExists checks if the directory exists and user has permission to read it
func DirExists(dir string) bool {
	info, err := os.Stat(dir)
	return !os.IsNotExist(err) && info.IsDir()
}

//ListFiles prints all the files for a given directory
func ListFiles(dir string) {
	fmt.Println(dir + ":")
	files, err := ioutil.ReadDir(".")
	if err == nil {
		for _, file := range files {
			fmt.Print(file.Name() + " ")
		}
	}
}

//GetDirsToList returns all valid and accessible dirs from a given list of dirs
func GetDirsToList() []string {
	argsWithoutProg := os.Args[1:]
	inputDirectories := []string{}
	if len(argsWithoutProg) > 0 {
		for _, currentDir := range argsWithoutProg {
			if ArgIsDir(currentDir) && DirExists(currentDir) {
				inputDirectories = append(inputDirectories, currentDir)
			}
		}
	} else {
		path, err := os.Getwd()
		if err == nil {
			inputDirectories = append(inputDirectories, path)
		}
	}
	return inputDirectories
}

func main() {
	directoriesToList := GetDirsToList()
	for _, currentDir := range directoriesToList {
		ListFiles(currentDir)
		fmt.Println()
	}
}

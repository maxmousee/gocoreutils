package main

import (
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
	return !os.IsNotExist(err) && !info.IsDir()
}

//ReadFile returns a string with file contents
func ReadFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(content)
}
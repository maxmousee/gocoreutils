package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoFilesToRead(t *testing.T) {
	assertions := assert.New(t)
	result := GetFilesToRead()
	assertions.Equal([]string{}, result)
}

func TestArgIsFile(t *testing.T) {
	assertions := assert.New(t)
	result := ArgIsFile("somefile")
	assertions.True(result, "strings not starting with dashes are recognized as files")
}

func TestArgIsNotAFile(t *testing.T) {
	assertions := assert.New(t)
	result := ArgIsFile("-x")
	assertions.False(result, "strings starting with dashes are not recognized as files")
}

func TestFileDoesNotExist(t *testing.T) {
	assertions := assert.New(t)
	result := FileExists("somefile.txt")
	assertions.False(result, "this file does not exist")
}

func TestFileExists(t *testing.T) {
	assertions := assert.New(t)
	result := FileExists("../testfiles/file1.txt")
	assertions.True(result, "this test file should exist")
}

func TestReadFile(t *testing.T) {
	assertions := assert.New(t)
	result := ReadFile("../testfiles/file1.txt")
	assertions.Equal("line1\nline2", result)
}

func TestReadFileThatDoesNotExist(t *testing.T) {
	assertions := assert.New(t)
	result := ReadFile("nonexistentfile.txt")
	assertions.Equal("", result)
}
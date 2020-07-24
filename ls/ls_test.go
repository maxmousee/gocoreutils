package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArgIsDir(t *testing.T) {
	assertions := assert.New(t)
	result := ArgIsDir("somedir")
	assertions.True(result, "strings not starting with dashes are recognized as dirs")
}

func TestArgIsNotADir(t *testing.T) {
	assertions := assert.New(t)
	result := ArgIsDir("-x")
	assertions.False(result, "strings starting with dashes are not recognized as dirs")
}

func TestNoDirsToRead(t *testing.T) {
	assertions := assert.New(t)
	result := GetDirsToList()
	assertions.Equal([]string{}, result)
}


func TestDirDoesNotExist(t *testing.T) {
	assertions := assert.New(t)
	result := DirExists("somedir")
	assertions.False(result, "this dir does not exist")
}

func TestDirExists(t *testing.T) {
	assertions := assert.New(t)
	result := DirExists("../testfiles/")
	assertions.True(result, "this test dir should exist")
}
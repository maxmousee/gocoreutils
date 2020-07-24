package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func TestFileExists(t *testing.T) {
	//TODO
}

func TestReadFile(t *testing.T) {
	//TODO
}
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_noFilesToRead(t *testing.T) {
	assertions := assert.New(t)
	result := getFilesToRead()
	assertions.Equal([]string{}, result)
}
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetTimeToSleepInSecs(t *testing.T) {
	assertions := assert.New(t)
	result := GetTimeToSleepInSecs("3")
	assertions.Equal(time.Duration(3)*time.Second, result)
}

func TestGetTimeToSleepInSecsReturnsZeroForInvalidString(t *testing.T) {
	assertions := assert.New(t)
	result := GetTimeToSleepInSecs("abc")
	assertions.Equal(time.Duration(0)*time.Second, result)
}
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

//GetTimeToSleepInSecs returns a time duration (in seconds) based on a given string
func GetTimeToSleepInSecs(inputArg string) time.Duration {
	timeInSecs, err := strconv.Atoi(inputArg)
	if err != nil || timeInSecs < 0 {
		return 0
	}
	return time.Duration(timeInSecs)*time.Second
}


func main() {
	if len(os.Args[1:]) == 1 {
		firstArgString := os.Args[1:][0]
		timeToSleep := GetTimeToSleepInSecs(firstArgString)
		time.Sleep(timeToSleep)
	} else {
		fmt.Println("usage: sleep seconds")
	}
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func getTimeToSleepInSecs(inputArg string) time.Duration {
	timeInSecs, err := strconv.Atoi(inputArg)
	if err != nil || timeInSecs < 0 {
		return 0
	}
	return time.Duration(timeInSecs)*time.Second
}


func main() {
	if len(os.Args[1:]) == 1 {
		firstArgString := os.Args[1:][0]
		timeToSleep := getTimeToSleepInSecs(firstArgString)
		time.Sleep(timeToSleep)
	} else {
		fmt.Println("usage: sleep seconds")
	}
}

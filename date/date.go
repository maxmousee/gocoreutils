package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	dateString := currentTime.Format("Mon Jan 02 15:04:05")
	dateString += " " + currentTime.Local().Format("CEST")
	dateString += " " + currentTime.Format("2006")
	fmt.Println(dateString)
}

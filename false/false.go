package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputArgs := os.Args[1:]
	inputArgsString := strings.Join(inputArgs, " ")
	if len(inputArgsString) > 0 {
		fmt.Println("false " + inputArgsString)
	} else {
		fmt.Println("false")
	}
}

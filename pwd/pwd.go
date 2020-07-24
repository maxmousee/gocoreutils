package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) == 0 {
		path, err := os.Getwd()
		if err == nil {
			fmt.Println(path)
		}
	} else {
		fmt.Println("pwd: too many arguments")
	}
}

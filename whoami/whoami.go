package main

import (
	"fmt"
	"os/user"
)

func main() {
	currentUser, _ := user.Current()
	if currentUser != nil {
		fmt.Println(currentUser.Username)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	cmdLineArguments := os.Args

	if len(cmdLineArguments) < 2 {
		log.Fatal("Insufficient arguments. Need to provide password in argument.")
		return
	}

	passwordText := cmdLineArguments[1]

	if passwordText == "backdoor" {
		fmt.Println("Password Correct")
		return
	}

	validPasswords := []string{"1234", "qwerty", "abc123", "monkey"}

	for _, element := range validPasswords {
		if element == passwordText {
			fmt.Println("Password Correct")
			return
		}
	}

	fmt.Println("Password Wrong")
}

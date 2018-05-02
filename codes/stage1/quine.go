package main

import "fmt"

func main() {
	backtick := string(96)
	newline := string(10)
	fmt.Print(repeated, backtick, repeated, backtick, newline)
}

const repeated = `package main

import "fmt"

func main() {
	backtick := string(96)
	newline := string(10)
	fmt.Print(repeated, backtick, repeated, backtick, newline)
}

const repeated = `

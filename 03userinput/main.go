package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Go programming!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your name: ")

	// comma ok || error ok syntax
	name, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Println("Hello, ", name)
	fmt.Printf("Type of name: %T\n", name)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome!")

	fmt.Println("Please rate our service (1-5): ")
	reader := bufio.NewReader(os.Stdin)

	rating, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Println("Thank you for your rating:", rating)
	fmt.Printf("Type of rating: %T\n", rating)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	if err != nil {
		fmt.Println("Error converting rating to number:", err)
		return
	}
	fmt.Printf("Your rating as a number: %.2f\n", numRating)
	fmt.Printf("Type of numRating: %T\n", numRating)
}

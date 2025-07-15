package main

import "fmt"

// First letter capitalized makes it public
const MyJwtToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6Ik1haW4gVXNlciIsImlhdCI6MTYxNjYwMDQwMH0.4"

func main() {
	var name string = "Muhammad Saqlain"
	fmt.Println("Hello, " + name + "!")
	fmt.Printf("Variable type: %T\n", name)

	var isVerified bool = true
	fmt.Println("Is Verified: ", isVerified)
	fmt.Printf("Variable type: %T\n", isVerified)

	var smallVal uint8 = 255
	fmt.Println("Small value: ", smallVal)
	fmt.Printf("Variable type: %T\n", smallVal)

	var smallFloat float64 = 255.2349854698793450987
	fmt.Println("Small float value: ", smallFloat)
	fmt.Printf("Variable type: %T\n", smallFloat)

	// default values and some aliases
	var emptyInt int
	fmt.Println("Empty int variable value: ", emptyInt)
	fmt.Printf("Variable type: %T\n", emptyInt)

	var emptyStr string
	fmt.Println("Empty string variable value: ", emptyStr)
	fmt.Printf("Variable type: %T\n", emptyStr)

	// implicit type
	var noDatatypeGiven = "Implicit type example"
	fmt.Println("No datatype variable value: ", noDatatypeGiven)
	fmt.Printf("Variable type: %T\n", noDatatypeGiven)

	// no var style (only allowed in function/method scope)
	noOfLines := 10
	fmt.Println("No var style variable value: ", noOfLines)
	fmt.Printf("Variable type: %T\n", noOfLines)

	fmt.Println("JWT Token: ", MyJwtToken)
	fmt.Printf("JWT Token type: %T\n", MyJwtToken)
}

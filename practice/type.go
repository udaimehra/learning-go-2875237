package main

import (
	"fmt"
)

var aConst int = 10

func main() {
	var aString string = "hello from go"
	fmt.Println(aString)
	fmt.Printf("The variable's type is %T \n", aString)
	fmt.Println("Hello from Go!")
	var anInteger int = 42
	fmt.Printf("The variable's type is %T \n", anInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	var anotherString string = "This is another string"
	fmt.Println(anotherString)
	fmt.Printf("The variable's type is %T \n", anotherString)

	var anotherInt int = 53
	fmt.Println(anotherInt)
	fmt.Printf("The variable's type is %T \n", anotherInt)

	myString := "This is also a string"
	fmt.Println(myString)
	fmt.Printf("The variable's type is %T \n", myString)

	fmt.Println(aConst)
	fmt.Printf("The type is %T \n", aConst)
	// This is a test line.
	// Adding another line.
}

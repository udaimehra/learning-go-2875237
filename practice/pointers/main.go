package main

import (
	"fmt"
)

func main() {

	//Start your program here.
	anInt := 42
	var p *int
	var p1 = &anInt
	fmt.Println("Value of p:", p)
	fmt.Println("Value of p1:", *p1)

	/* Following example shows how the value of a variable can be changed
	*/ by doing an opertion on the pointer on that object.
	aFloat := 10.4
	var value1 = &aFloat
	fmt.Println("Value of float is: ", *value1)
	*value1 = *value1 / 8
	fmt.Println("Value of value1 after division with 8 is: ", *value1)
	fmt.Println("Value of aFloat after division with 8 is: ", aFloat)
}

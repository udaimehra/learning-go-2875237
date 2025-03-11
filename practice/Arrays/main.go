package main

import (
	"fmt"
)

func main() {

	// Write your code below.
	fmt.Println("---Arrays---")
	var colors [3]string
	var colorSlice = []string{"Red", "Blue", "Green"}
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println("Array is: ", colors)
	fmt.Println("Slice of array is: ", colorSlice)

	var numbers = [5]int{1, 2, 3, 4, 5}
	/*	numbers[0] = 1
		numbers[1] = 2
		numbers[2] = 3
		numbers[3] = 4
		numbers[4] = 5
	*/
	fmt.Println("Array of numbers is:", numbers)
	fmt.Println("What is the length of colors array? ", len(colors))
	fmt.Println("What is the length of numbers array? ", len(numbers))

}

package main

import (
	"fmt"
	"math"
)

var pi float64 = 3.14159

func main() {
	i1, i2, i3 := 12, 14, 68
	intSum := i1 + i2 + i3
	fmt.Println("Sum of integers is: ", intSum)
	f1, f2, f3 := 1.13, 1.29, 1.45
	fltSum := f1 + f2 + f3
	fltSum = math.Round(fltSum*100) / 100
	fmt.Println("Rounded sum of floats is: ", fltSum, "\n")
	fmt.Printf("Rounded sum of floats is: %v\n", fltSum)
	rpi := math.Round(pi)
	fmt.Printf("Original value of pi is: %v and Rounded value of pi is: %v \n", pi, rpi)
	fmt.Printf("Rounded value of pi is: %v\n", rpi)

	circleRadius := 5.2
	circumference := 2 * math.Pi * circleRadius
	fmt.Printf("Circumferene is: %.2f\n", circumference)

}

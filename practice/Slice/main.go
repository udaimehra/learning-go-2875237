package main

import (
	"fmt"
	"sort"
)

func main() {
	// Write your program below.
	fmt.Println("Slices in Arrays")
	var slice = []string{"Red", "Green", "Blue"}
	fmt.Println("The full slice is: ", slice)
	// Adding / Appending another item at the end of the slice.
	slice = append(slice, "Purple")
	fmt.Println("Slice after adding another item is now:", slice)
	//Print the slice with first item eliminated.
	slice = append(slice[1:len(slice)])
	fmt.Println("After eliminating first item:", slice)
	//Print the slice after eliminating last item, but including first item.
	slice = append(slice[:len(slice)-1])
	fmt.Println("After eliminating last item from the slice", slice)
	fmt.Println("The state of slice after all these operations: ", slice)
	fmt.Println("Adding another item called Magenta in it")
	slice = append(slice, "Magenta")
	fmt.Println("After adding Magenta:", slice)

	// Learn the use of make for slices
	slice1 := make([]int, 5, 10)
	slice1[0] = 500
	slice1[1] = 115
	slice1[2] = 2
	slice1[3] = 35
	slice1[4] = 425
	fmt.Println("Slice created with the help of make(): ", slice1)
	sort.Ints(slice1)
	fmt.Println(slice1)

}

package main

import (
 	"fmt"
 	"strconv"
 	"strings"
 )

func sum(value1 string, value2 string) float64 {
	fltValue1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
        if err != nil {
	fmt.Println(err)
	panic("Value 1 must be a number")
        }
	fltValue2, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)
        if err != nil {
	fmt.Println(err)
	panic("Value 2 must be a number")
        }
	return fltValue1 + fltValue2
}
var value1 string = "10"
var value2 string = "5.5"	
var result float64
result := sum(value1, value2)
fmt.Printf("Sum of these numbers is %.2f\n", result)

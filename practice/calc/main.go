// Write your answer here, and then test your code.
package main

// Uncomment this import section to use required Go packages
import (
 	"fmt"
 	"strconv"
 	"strings"
 	"bufio"
 	"os"
)

// sum() returns the sum of the two parameters
func sum(value1 string, value2 string) float64 {
	fltValue1, _ := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	fltValue2, _ := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	result := fltValue1 + fltValue2
	return result
}
// prod() returns the product of two parameters
func prod(value1 string, value2 string) float64 {
	fltValue1, _ := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	fltValue2, _ := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	result := fltValue1 * fltValue2
	return result
}

func main() {
	// Take first string as input from the user
	reader1 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter first number: ")
        value1, _ := reader1.ReadString('\n')
	fmt.Println("You entered", value1)

	// Take first string as input from the user
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter first number: ")
        value2, _ := reader2.ReadString('\n')
	fmt.Println("You entered", value2)
	fmt.Printf("Sum of these numbers is %.2f\n", sum(value1, value2))
	fmt.Printf("Product of these numbers is %.2f\n", prod(value1, value2))
}

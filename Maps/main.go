package main

import (
	"fmt"
)

func main() {

	// Write your code below this line.
	fmt.Println("Maps	")
	states := make(map[string]string)
	fmt.Println(states)
	states["UP"] = "Uttar Pradesh"
	states["HP"] = "Himachal Pradesh"
	states["MP"] = "Madhya Pradesh"
	fmt.Println(states)
	fmt.Println("The abbreviaion of MP, stands for: ", states["MP"])
	uttarakhand := states["UK"]
	fmt.Println(states)
}

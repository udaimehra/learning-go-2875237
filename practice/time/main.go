package main

import (
	"fmt"
	"time"
)

func main() {
	old := time.Now()
	fmt.Println("Dates and Times")
	fmt.Println(old.Format(time.ANSIC))
	fmt.Println(old)
}

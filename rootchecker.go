package main 

import (
	"fmt"
)

func main() {
	num := 0
	fmt.Print("enter number: ")
	fmt.Scan(&num)

	fmt.Println("the square root of", num, "is", num*num)
}
package main

import "fmt"

func main() {
	var length int
	var width int

	fmt.Print("input your length: ")
	fmt.Scan(&length)

	fmt.Print("input your width: ")
	fmt.Scan(&width)

	result := width * length

	fmt.Printf("the lenght is %v and the width is %v \n", length, width)
	fmt.Println("the final output is:", result)
}

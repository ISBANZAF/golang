package main 

import "fmt"

func main() {
	var age int 

	fmt.Print("input your age: ")

	fmt.Scan(&age)

	if age > 18 {
		fmt.Println("successfully registered!")
	} else if age < 18 {
		fmt.Println("opps, you've got ways to go")
	}
}
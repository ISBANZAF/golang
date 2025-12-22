package main

import "fmt"

func CreateGreeting(name string) string {
	return "hello " + name + " welcome!"
}

func main() {
	var myname string
	fmt.Print("input name: ")
	fmt.Scan(&myname)

	result := CreateGreeting(myname)
	fmt.Println(result)
}

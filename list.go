package main

import "fmt"

func main() {
	items := []string{"apple", "milk", "sugar"}

	items = append(items, "butter")

	fmt.Println("my shop list are: ")
	for i, item := range items {
		fmt.Printf("%d, %s\n", i+1, item)
	}
}

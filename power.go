package main

import (
	"fmt"
	"time"
)

func count(s string) {
	for i := 1; i < 5; i++ {

		fmt.Printf("%s: %d\n", s, i)
		time.Sleep(time.Second * 5)
	}
}

func main() {
	go count("Background Task")
	count("Main Task")

	fmt.Println("Done!")
}

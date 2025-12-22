package main

import (
	"fmt"
	"time"
)

func Quad(x, y int) {

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if i == 0 && j == 0 {
				fmt.Print("/")
			} else if i == 0 && j == y-1 {
				fmt.Print("\\")
				time.Sleep(time.Second * 1)
			} else if i == x-1 && j == y-1 {
				fmt.Print("/")
				time.Sleep(time.Second * 1)
			} else if i == x-1 && j == 0 {
				fmt.Print("\\")
				time.Sleep(time.Second * 1)
			} else if i == 0 || j == 0 || j == y-1 || i == x-1 {
				if i == 0 && j == 6 || i == x-1 && j == 6 {
					continue
				}
				fmt.Print("*")
				time.Sleep(time.Second * 1)
			} else {
				fmt.Print(" ")
			} 
		} 
		fmt.Println()
	}
}


func main() {
	Quad(4, 6)
}
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i <= 10; i ++ {
		fmt.Println("the suquare root of", i, "is", i*i)
		time.Sleep(time.Second * 1)
	}
}
package main

import (
	"fmt"
	"time"
)

func main() {
	password := ""

	fmt.Print("Create Password: ")
	fmt.Scan(&password)
	time.Sleep(time.Second * 2)
	if len(password) < 8 {
		fmt.Println("password must be 8 or more!")
		fmt.Print("enter password again: ")
		fmt.Scan(&password)
		time.Sleep(time.Second * 2)
	}
	fmt.Printf("congratulation youve created your first password:: %s\n", password)
}

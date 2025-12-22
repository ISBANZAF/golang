package main

import "fmt"

type user struct {
	name   string
	age    int
	Active bool
}

func main() {
	player := user{
		name:   "",
		age:    21,
		Active: true,
	}
	fmt.Print("input name: ")
	fmt.Scan(&player.name)
	fmt.Printf("congrats on signing up %s\n", player.name)
	fmt.Printf("you're qualified to access our game at the age of %d\n", player.age)

	if player.Active {
		fmt.Println("status: Online")
	}
}

package main

import "fmt"

type character struct {
	name   string
	level  int
	health float64
	hero   bool
}

func main() {
	player := character{
		name:   "",
		level:  0,
		health: 100.0,
		hero:   true,
	}

	fmt.Print("name yourself: ")
	fmt.Scan(&player.name)

	fmt.Print("start from any level you wish: ")
	fmt.Scan(&player.level)

	fmt.Printf("your name is %s\n", player.name)
	fmt.Printf("your level is %d\n", player.level)
	fmt.Printf("your health is now %.1f\n", player.health)
	fmt.Printf("are you a hero thou? %v\n", player.hero)
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	secret := rand.Intn(10) + 1

	var guess int
	fmt.Print("Guess the number (1–10): ")
	fmt.Scanln(&guess)

	for {
		if guess != secret {

			if guess < secret {
		fmt.Println("the answer is greater than that!")
	} else if guess > secret {
		fmt.Println("the answer is less than that!")
	}
			fmt.Println("try again: ")
			fmt.Scanln(&guess)
		}
		if guess == secret {
			fmt.Println("Congratulatons The Number is:", guess)
				break
		}
	}
}

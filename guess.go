package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	secret := rand.Intn(100) + 1

	var guess int
	fmt.Print("Guess the number (1 to 100): ")
	fmt.Scanln(&guess)

	chance := 5

	for chance > 0 {
		if guess != secret {

			if guess < secret {
		fmt.Println("the answer is greater than that!")
	} else if guess > secret {
		fmt.Println("the answer is less than that!")
	}
			fmt.Println("try again: ")
			fmt.Scanln(&guess)
		}

		chance --
		fmt.Printf("you have %d chance left!\n", chance)


		if guess == secret {
			fmt.Println("Congratulatons The Number is:", guess, "\n")

			fmt.Println("moving to the next level!\n this will now add 50 to the former maximum number")
			time.Sleep(time.Second * 2)

			secret += 50 
			chance += 2

			fmt.Println("guess the number: ")
			fmt.Scanln(&guess)

				if guess != secret {

			if guess < secret {
		fmt.Println("the answer is greater than that!")
	} else if guess > secret {
		fmt.Println("the answer is less than that!")
	}
			fmt.Println("try again: ")
			fmt.Scanln(&guess)
		}

		chance --
		fmt.Printf("you have %d chance left!\n", chance)


		if guess == secret {
			fmt.Println("Congratulatons The Number is:", guess, "\n")
			}

		}
	}
}

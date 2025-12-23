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
	fmt.Println("LEVEL 1")
	fmt.Print("Guess the number (1 to 100): ")
	fmt.Scanln(&guess)
	time.Sleep(time.Second * 2)

	level := 1
	chance := 5
	levelnum := 100
	for chance > 0 {
		if guess != secret {

			if guess < secret {
		fmt.Println("the answer is greater than that!")
	} else if guess > secret {
		fmt.Println("the answer is less than that!")
	}
			fmt.Println("try again: ")
			fmt.Scanln(&guess)
			time.Sleep(time.Second * 1)
		}

		chance --
		fmt.Printf("you have %d chance left!\n", chance)


		if guess == secret {
			fmt.Println("Congratulatons The Number is:", guess, "\n")
			level ++
			levelnum += 50
			fmt.Printf("moving to LEVEL %d!\nwhich will be %d ", level, levelnum)
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
			time.Sleep(time.Second * 1)
		}

		chance --
		fmt.Printf("you have %d chance left!\n", chance)


		if guess == secret {
			fmt.Println("Congratulatons The Number is:", guess, "\n")
			}

		}
	}
}

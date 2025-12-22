package main

import "fmt"

func main() {
	var kms float64

	const milesperkm = 0.621371

	fmt.Print("enter distance in kilometers: ")
	fmt.Scan(&kms)

	miles := kms * milesperkm 

	fmt.Printf("%.2f kilometer is roughly %.2f miles \n", kms, miles)
}
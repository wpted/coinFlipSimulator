package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Enter how many times you want to flip the coin:")
	var count int
	// Assign the user input to the address of count: &count
	fmt.Scanln(&count)
	tails := 0
	heads := 0
	var results []string

	for i := 0; i < count; i++ {
		result := headOrTails()
		results = append(results, result)
		if result == "Head" {
			heads++
		} else {
			tails++
		}
	}
	fmt.Println("Coin_Flip result is,", results)
	fmt.Printf("Flip %d times. Heads %d times, tails %d times.", count, heads, tails)

}

func headOrTails() string {
	// Seeding with something like the current time gives the random function a different number everytime
	rand.Seed(time.Now().UnixNano())
	// 0 <= r < 2
	r := rand.Intn(2)

	if r == 0 {
		return "Head"
	} else {
		return "Tails"
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Here I will compile this task
// Write a program that will guess this number.
// Make the computer pick random numbers from the gap until it guesses the number you have in mind.
// This number must be announced in advance at the top of the program.
// Let each assumption be displayed on the screen with an explanation of whether this number is greater
// or less than the option you have conceived

func guessTheNumber(number int) {
	rand.Seed(time.Now().UnixNano())
	for {
		var estimatedNumber = rand.Intn(number) + 1
		println(estimatedNumber)
		if estimatedNumber == number {
			fmt.Println("You guessed it!")
			break
		}
		if estimatedNumber < number {
			fmt.Println("You guessed wrong. My number is bigger")
		}
		if estimatedNumber > number {
			fmt.Println("You guessed wrong. My number is smaller")
		}
	}
}

func main() {
	guessTheNumber(200)
}

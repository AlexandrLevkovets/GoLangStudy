package main

import "fmt"

// Here I will compile this task
// There are several locations in the adventure game in question.
// Write a program that uses if to
// display a description of each of the three locations: caves, entrances, and mountains

func displayDescription(location string) {
	if location == "caves" {
		fmt.Println("You don't see anything")
	}
	if location == "entrances" {
		fmt.Println("You see dark cave")
	}
	if location == "mountains" {
		fmt.Println("You see beautiful landscape")
	}
}

func main() {
	displayDescription("caves")
	displayDescription("entrances")
	displayDescription("mountains")
}

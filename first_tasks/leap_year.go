package main

import "fmt"

// Here I will compile this task
//Find out if the year was a leap year.

func isItLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}

func main() {
	var year = 2000
	fmt.Printf("Is it %v leap year? Answer in bool: %v\n", year, isItLeapYear(year))
	year = 2002
	fmt.Printf("Is it %v leap year? Answer in bool: %v\n", year, isItLeapYear(year))
	year = 2100
	fmt.Printf("Is it %v leap year? Answer in bool: %v\n", year, isItLeapYear(year))
	year = 2132
	fmt.Printf("Is it %v leap year? Answer in bool: %v\n", year, isItLeapYear(year))
}

package main

import (
	"fmt"
	"math"
)

// Here I will compile this task
// https://leetcode.com/problems/palindrome-number/description/

// This algo was created without any analyzing... just for beat my self
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	length := (int)(math.Log10(float64(x)) + 1)
	if length <= 1 {
		return true
	}
	var NumberArray []int
	for counter := length; counter >= 1; counter-- {
		_, afterPoint := math.Modf(float64(x) / math.Pow(10, float64(counter)))
		afterPoint = math.Round(afterPoint * math.Pow(10, float64(1)))
		NumberArray = append(NumberArray, int(afterPoint))
	}
	var reverseNumberArray []int
	for counter := 1; counter <= length; counter++ {
		_, afterPoint := math.Modf(float64(x) / math.Pow(10, float64(counter)))
		if counter == 1 {
			if afterPoint >= float64(NumberArray[len(NumberArray)-1])/10 {
				afterPoint = math.Floor(afterPoint * math.Pow(10, float64(1)))
			} else {
				afterPoint = math.Ceil(afterPoint * math.Pow(10, float64(1)))
			}
		} else {
			afterPoint = math.Floor(afterPoint * math.Pow(10, float64(1)))
		}
		reverseNumberArray = append(reverseNumberArray, int(afterPoint))
	}
	reverseNumber := 0
	for counter := 0; counter <= len(reverseNumberArray)-1; counter++ {
		reverseNumber += int(float64(reverseNumberArray[counter]) * math.Pow(10, float64(length-counter-1)))
	}
	if reverseNumber == x {
		return true
	}
	return false
}

// this was created after 20 minutes of thinking XD and read about % operation in Go...
func isPalindrome2(x int) bool {
	if (x != 0 && x%10 == 0) || x < 0 {
		return false
	}
	reversedNumber := 0
	for reversedNumber < x {
		reversedNumber = reversedNumber*10 + x%10
		x = x / 10
	}
	return x == reversedNumber/10 || x == reversedNumber
}

func main() {
	fmt.Println(isPalindrome(11))
	fmt.Println(isPalindrome2(121))
}

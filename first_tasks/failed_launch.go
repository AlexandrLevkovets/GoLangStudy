package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Here I will compile this task
// P.S according the task I must use simple version of for
// Not every launch goes according to plan.
// Implement a countdown, where for every second there is a 1 in 100 chance that due to certain
// circumstances the launch will be interrupted and the counter will stop.

func startLaunch(secondsToLaunch int) bool {
	for secondsToLaunch != 0 {
		time.Sleep(time.Second)
		if (rand.Intn(100) + 1) == 100 {
			return false
		}
		secondsToLaunch--
	}
	return true
}

func main() {
	fmt.Printf("Is the launch successful? Answer in bool: %v\n", startLaunch(10))
}

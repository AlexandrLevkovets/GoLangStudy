package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Here I will compile this task
// SpaceX's interplanetary transport system clearly lacks a warp drive, but it can reach Mars at a speed of 100,800 km/h.
// The trip is scheduled for January 2025. The distance between Mars and Earth at this time will be 96,300,000 km.
// How many days will astronauts fly to Mars?

// As part of the task, I decided that if we count the number of days,
// it would be better to round them up, it seems to me more correct

func travelFromEarthToMars() {
	const hoursInDay = 24
	var speedOfShip = 100800.0        // km/h
	var distanceOfTravel = 96900000.0 // km
	fmt.Printf("This trip will take this count of days %v\n", math.Ceil(distanceOfTravel/(speedOfShip*hoursInDay)))
}

// Also, as part of the task, I decided to add a new condition:

// The distance between Earth and Mars differs at different times and depends on where the planets are
// orbiting the Sun at a given time.
// Add the condition that the distance becomes random from 56,000,000 to 401,000,000 km.

func travelFromEarthToMarsCalculated(distanceOfTravel int) {
	const hoursInDay = 24
	var speedOfShip = 100800.0
	fmt.Printf("This trip will take this count of days %v\n", math.Ceil(float64(distanceOfTravel)/(speedOfShip*hoursInDay)))
}

func main() {
	travelFromEarthToMars()
	fmt.Println(rand.Intn(401000001-56000000) + 56000000)
	travelFromEarthToMarsCalculated(rand.Intn(401000001-56000000) + 56000000)
}

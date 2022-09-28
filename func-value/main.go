package main

import (
	"fmt"
	"math"
)

func main() {
	// teacher mentioned not using single letter variables ;)
	fmt.Printf("Enter acceleration: ")
	var accel float64
	fmt.Scan(&accel)

	fmt.Printf("Enter initial velocity: ")
	var initVelo float64
	fmt.Scan(&initVelo)

	fmt.Printf("Enter initial displacement: ")
	var initDisplace float64
	fmt.Scan(&initDisplace)

	fmt.Printf("Enter time: ")
	var time float64
	fmt.Scan(&time)

	displaceFn := GenDisplaceFn(accel, initVelo, initDisplace)

	fmt.Printf("Displacement: %v\n", displaceFn(time))
}

func GenDisplaceFn(accel, initVelo, initDisplace float64) func(t float64) float64 {
	displaceFn := func(time float64) float64 {
		displace := 0.5*accel*math.Pow(time, 2) + initVelo*time + initDisplace
		return displace
	}

	return displaceFn
}

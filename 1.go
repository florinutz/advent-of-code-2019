package main

import (
	"fmt"
	"math"
)

func main() {
	modules := []int64{
		1969,
	}

	var sum int64 = 0

	for _, mass := range modules {
		sum += getFuel(mass)
	}

	fmt.Println(sum)
}

func getFuel(mass int64) int64 {
	fuel := int64(math.Floor(float64(mass)/3)) - 2

	if fuel <= 0 {
		return 0
	}

	return fuel + getFuel(fuel)
}

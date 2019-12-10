package main

import (
	"fmt"
	"log"
)

func main() {
	a := []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 13, 19, 1, 9, 19, 23, 1, 6, 23, 27, 2, 27, 9, 31, 2, 6, 31, 35, 1, 5, 35, 39, 1, 10, 39, 43, 1, 43, 13, 47, 1, 47, 9, 51, 1, 51, 9, 55, 1, 55, 9, 59, 2, 9, 59, 63, 2, 9, 63, 67, 1, 5, 67, 71, 2, 13, 71, 75, 1, 6, 75, 79, 1, 10, 79, 83, 2, 6, 83, 87, 1, 87, 5, 91, 1, 91, 9, 95, 1, 95, 10, 99, 2, 9, 99, 103, 1, 5, 103, 107, 1, 5, 107, 111, 2, 111, 10, 115, 1, 6, 115, 119, 2, 10, 119, 123, 1, 6, 123, 127, 1, 127, 5, 131, 2, 9, 131, 135, 1, 5, 135, 139, 1, 139, 10, 143, 1, 143, 2, 147, 1, 147, 5, 0, 99, 2, 0, 14, 0}

	result := day2_1(a, 12, 2)
	fmt.Printf("Day 2.1 result: %d\n", result)

	day22WantedValue := 19690720
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			if day2_1(a, noun, verb) == day22WantedValue {
				fmt.Printf("%d%d\n", noun, verb)
			}
		}
	}
}

var ops = map[int]func(a []int, currentAddr int) []int{
	1: func(a []int, currentAddr int) []int {
		dst := a[currentAddr+3]
		p1 := a[currentAddr+1]
		p2 := a[currentAddr+2]

		a[dst] = a[p1] + a[p2]

		return a
	},
	2: func(a []int, currentAddr int) []int {
		dst := a[currentAddr+3]
		p1 := a[currentAddr+1]
		p2 := a[currentAddr+2]

		a[dst] = a[p1] * a[p2]

		return a
	},
}

func day2_1(a []int, noun, verb int) int {
	b := make([]int, len(a))
	copy(b, a)

	b[1] = noun
	b[2] = verb

	i := 0
	for i <= len(b)-1 {
		if b[i] == 99 {
			break
		}
		if op, ok := ops[b[i]]; ok {
			b = op(b, i)
		} else {
			log.Fatalf("invalid op %d", b[i])
		}
		i += 4
	}

	return b[0]
}

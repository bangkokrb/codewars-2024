package main

import (
	"fmt"
	"math"
)

func sumInts(ints []int) int {
	s := 0
	for _, n := range ints {
		s += n
	}
	return s
}

func main() {
	ints := []int{-54, 39, -2, -88, 42, -18, -16, -16, 0, 22, 70, 55, -57, 43,
		-27, 88, 28, 6, 60, -39, -85, 46, -57, 83, 0, -53, 0, 10, 22, -78, 26, -7,
		100, -87, 47, 72, 94, -11, -42, 100, 63, -35, 39, 2, 57, -30, -17, -75, 27,
		83}

	maxS := math.MinInt
	for i := range ints {
		fmt.Printf("Considering %d\n", ints[i])
		for j := i + 1; j < len(ints)+1; j++ {
			s := sumInts(ints[i:j])
			if s > maxS {
				fmt.Printf("new max: %d ", s)
				fmt.Println(ints[i:j])
				maxS = s
			}
		}
	}

	fmt.Println(maxS)
	// new max: 555 [0 22 70 55 -57 43 -27 88 28 6 60 -39 -85 46 -57 83 0 -53 0 10 22 -78 26 -7 100 -87 47 72 94 -11 -42 100 63 -35 39 2 57]
}

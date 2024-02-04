package main

import (
	"fmt"
)

func isEven(n int) bool {
	return n%2 == 0
}

func main() {

	// conditions where you make a U-turn:
	// - odd soi: larger to smaller (2x)
	// - even soi: smaller to larger (2x)
	// - odd soi to even soi, or vice versa

	route := []int{2, 8, 10, 4, 6, 2, 10, 3, 11, 7, 8, 11, 15, 12, 5, 7, 14, 2,
		7, 14, 13, 7, 5, 12, 15, 11, 4, 11, 4, 11, 10, 5, 8, 2, 14, 2, 3, 10, 9, 11,
		6, 14, 1, 14, 12, 2, 9, 15, 8, 7, 3, 14, 8, 2, 12, 7, 5, 3, 10, 9, 12, 3, 8,
		11, 2, 4, 10, 6, 13, 15, 4, 5, 13, 3, 6, 3, 8, 7, 4, 1, 4, 8, 6, 5, 3, 14,
		10, 2, 6, 9, 3, 11, 12, 2, 1, 15, 9, 4, 13, 2}

	uturns := 0

	for i, nxt := range route {
		if i == 0 {
			continue
		}

		prev := route[i-1]

		if isEven(prev) && !isEven(nxt) {
			uturns += 1
			fmt.Printf("%02d >> %02d: U-TURN (even >> odd) [%02d]\n", prev, nxt, uturns)
			continue
		}

		if !isEven(prev) && isEven(nxt) {
			uturns += 1
			fmt.Printf("%02d >> %02d: U-TURN (odd >> even) [%02d]\n", prev, nxt, uturns)
			continue
		}

		if isEven(prev) && (prev < nxt) {
			uturns += 2
			fmt.Printf("%02d >> %02d: U-TURNx2 (even S>>L) [%02d]\n", prev, nxt, uturns)
			continue
		}

		if !isEven(prev) && (prev > nxt) {
			uturns += 2
			fmt.Printf("%02d >> %02d: U-TURNx2 (odd L>>S)  [%02d]\n", prev, nxt, uturns)
			continue
		}

		fmt.Printf("%02d >> %02d: no U-turn            [%02d]\n", prev, nxt, uturns)
	}

	fmt.Println(uturns)
}

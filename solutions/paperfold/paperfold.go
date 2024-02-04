package main

import (
	"fmt"
)

func main() {
	a := 210.0
	b := 297.0
	limit := 200
	nFolds := 0

	for i := 0; i < limit; i++ {
		if a > b {
			a = a / 2
		} else {
			b = b / 2
		}

		nFolds += 1

		if (a <= 2) && (b <= 2) {
			break
		}
	}

	fmt.Printf("Final size after %d folds: %f x %f\n", nFolds, a, b)
}

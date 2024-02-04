/*
# Coins

* A 10 baht coin weighs 8.5 grams
* A 5 baht coin weights 6 grams
* A 2 baht coin weighs 4 grams
* A 1 baht coin weighs 3 grams

If i have 9 grams of coins, I could have one 1 baht coin and one 5 baht coin, or
three 1 baht coins. That's 2 possible combinations.

If I have 55.5 grams of coins, how many different possible combinations of 1, 2,
5 and 10 baht coins could I have?

ANSWER: 30
*/

package main

import (
	"fmt"
)

func main() {
	// brute force solution
	var combinations [][]int

	for b10 := 0; b10 < 10; b10++ {
		for b5 := 0; b5 < 15; b5++ {
			for b2 := 0; b2 < 20; b2++ {
				for b1 := 0; b1 < 25; b1++ {
					wgt := (b10 * 85) + (b5 * 60) + (b2 * 40) + (b1 * 30)

					if wgt != 555 {
						continue
					}

					combinations = append(combinations, []int{b10, b5, b2, b1})
				}
			}
		}
	}

	for _, combo := range combinations {
		fmt.Println(combo)
	}

	fmt.Println(len(combinations))
}

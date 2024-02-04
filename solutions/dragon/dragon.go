package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	years := []int{2024}
	for i := 0; i < (4000/12 + 1); i++ {
		years = append(years, years[len(years)-1]+12)
	}

	var matches []int
	for _, yr := range years {
    if yr > 4000 {
      continue
    }

		sum := 0
		for _, ch := range strings.Split(strconv.Itoa(yr), "") {
			n, _ := strconv.Atoi(ch)
			sum += n
		}
		if sum == 8 {
			matches = append(matches, yr)
		}
	}

	fmt.Println(matches)
	fmt.Println(len(matches))
}

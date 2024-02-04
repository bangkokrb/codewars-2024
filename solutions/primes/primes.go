package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func validateArgs() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [target number]\n", os.Args[0])
		os.Exit(1)
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	max := (int)(math.Floor(math.Sqrt(float64(n))))

	for i := 2; i <= max; i++ {

		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	validateArgs()

	n, _ := strconv.Atoi(os.Args[1])

	// we will solve this the "hard" way:
	// - simple Prime test: check if number n is divisible by 2 >> sqrt(n);
	//   avoid using library Prime check
	// - iterate up and down from n;
	//   avoid using library Prime Sieve

	var primes []int

	for i := 0; len(primes) < 3; i++ {
		// put a limit on this loop
		if i > 100 && i > (n*10) {
      fmt.Println(primes)
			fmt.Printf("Exceeded max checks (%d)", i)
      break
		}

		checkPos := n + i
		checkNeg := n - i

		if isPrime(checkPos) {
			primes = append(primes, checkPos)
		}
		if (checkPos != checkNeg) && isPrime(checkNeg) {
			primes = append(primes, checkNeg)
		}
	}

	fmt.Println(primes)
  sum := 0
  for _, n := range primes {
    sum += n
  }
	fmt.Println(sum)
}

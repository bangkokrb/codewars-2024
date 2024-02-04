/*
# Flag Scaler

The independent nation of Mattistan has decided their flag will be made from
ASCII characters. Specifically, the flag looks like this:

+-+-+
|\|/|
+-+-+
|/|\|
+-+-+

This flag can be scaled to an arbitrary size N. A size 2 flag looks like this:

+--+--+
|\ | /|
| \|/ |
+--+--+
| /|\ |
|/ | \|
+--+--+

Write a program takes any size N and outputs a flag of that size.

*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validateArgs() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [size of flag]\n", os.Args[0])
		os.Exit(1)
	}
}

func main() {
	validateArgs()

	n, _ := strconv.Atoi(os.Args[1])

	var bar string

	{
		var sb strings.Builder
		sb.WriteString("+")
		for i := 0; i < n; i++ {
			sb.WriteString("-")
		}
		sb.WriteString("+")
		for i := 0; i < n; i++ {
			sb.WriteString("-")
		}
		sb.WriteString("+")
		bar = sb.String()
	}

	var upper string

	{
		var sb strings.Builder

		for i := 0; i < n; i++ {
			sb.WriteString("|")

			for j := 0; j < n; j++ {
				if i == j {
					sb.WriteString(`\`)
				} else {
					sb.WriteString(" ")
				}
			}

			sb.WriteString("|")

			for j := 0; j < n; j++ {
				if i == (n - j - 1) {
					sb.WriteString(`/`)
				} else {
					sb.WriteString(" ")
				}
			}

			sb.WriteString("|")

      if i < n - 1 {
        sb.WriteString("\n")
      }
		}

		upper = sb.String()
	}

	var lower string

	{
		var sb strings.Builder

		for i := 0; i < n; i++ {
			sb.WriteString("|")

			for j := 0; j < n; j++ {
				if i == (n - j - 1) {
					sb.WriteString(`/`)
				} else {
					sb.WriteString(" ")
				}
			}

			sb.WriteString("|")

			for j := 0; j < n; j++ {
				if i == j {
					sb.WriteString(`\`)
				} else {
					sb.WriteString(" ")
				}
			}

			sb.WriteString("|")

      if i < n - 1 {
        sb.WriteString("\n")
      }
		}

		lower = sb.String()
	}

	fmt.Println(bar)
	fmt.Println(upper)
	fmt.Println(bar)
	fmt.Println(lower)
	fmt.Println(bar)

}

/*
---
answer: 40
---
# Word Search

A word search puzzle is a grid of letters where your challenge is to find
selected words as formed by consecutive letters in a line along the rows or
columns of the grid. In this puzzle, a word can only appear forwards and we
will ignore diagonals.

For example, the grid below:

```
DOG
WLE
PDT
```

contains the words DOG, OLD, and GET.

Find all the words _4 letters_ or greater in the word search below. A
dictionary.txt file is provided.

* Output each unique word found
* Provide the **number of words found** as your answer

```
CYSCNRLGUNXOFMCBKLBWAWAUC
VNXWHMEXDZJXZSTUYCGDPBNZZ
RQARFLEBQMYTWAZBGUXYRXGMU
HAHZFFAKXJALYEPRLSNJBXDAL
UUAMOWHHYSPXHFTUUCIDFCAPV
ACKIWDGAUAEVHEOPPIPZGSHOA
VLCGJWYSZLUSKFMFKLPYTHONH
GIEBKGIKLTYZZADKLIXRRIRHC
XQJFLIVETSIQPYGCISGUDWWHY
REPTCHJLYZEUTYUTKPTBSWLKS
NKDEIMCLWEMMHVKECDOYIFEZC
JBMFNZILLJLJXZLYWIIGDWSCH
IWPVRBHQFCJHULHFSUAIHBFWE
DENAOCKJITXMONGLASSEMBLYM
MWHYUKXTLHXUDXZACRFOSYMJE
MEHIYHAEDHZXHUNUVLBYZYLCN
WDAZEZCZAUNTVNQDGPFTLDPFD
YOVDDHENCFPGWDTGEUCSHNUHB
DYHJKNXVCFTYPESCRIPTTTAOL
WMJRNDWFQKHJAVASCRIPTAQCC
CLOJURERXUKYOOFTKZWSJOPAW
EKDIPLLAJPJULIARKBUYPHIMQ
YWBCOOQSXIUENXLUYATDXATLZ
YPPLHSPJWEDSVELSSHESZQWKY
UUMFDWTSMTIMNZTTVMNHDCGNQ
```
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const wordSearch = `CYSCNRLGUNXOFMCBKLBWAWAUC
VNXWHMEXDZJXZSTUYCGDPBNZZ
RQARFLEBQMYTWAZBGUXYRXGMU
HAHZFFAKXJALYEPRLSNJBXDAL
UUAMOWHHYSPXHFTUUCIDFCAPV
ACKIWDGAUAEVHEOPPIPZGSHOA
VLCGJWYSZLUSKFMFKLPYTHONH
GIEBKGIKLTYZZADKLIXRRIRHC
XQJFLIVETSIQPYGCISGUDWWHY
REPTCHJLYZEUTYUTKPTBSWLKS
NKDEIMCLWEMMHVKECDOYIFEZC
JBMFNZILLJLJXZLYWIIGDWSCH
IWPVRBHQFCJHULHFSUAIHBFWE
DENAOCKJITXMONGLASSEMBLYM
MWHYUKXTLHXUDXZACRFOSYMJE
MEHIYHAEDHZXHUNUVLBYZYLCN
WDAZEZCZAUNTVNQDGPFTLDPFD
YOVDDHENCFPGWDTGEUCSHNUHB
DYHJKNXVCFTYPESCRIPTTTAOL
WMJRNDWFQKHJAVASCRIPTAQCC
CLOJURERXUKYOOFTKZWSJOPAW
EKDIPLLAJPJULIARKBUYPHIMQ
YWBCOOQSXIUENXLUYATDXATLZ
YPPLHSPJWEDSVELSSHESZQWKY
UUMFDWTSMTIMNZTTVMNHDCGNQ`

func validateArgs() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [dictionary file]\n", os.Args[0])
		os.Exit(1)
	}
}

func main() {
  validateArgs()

	// load and make the dictionary uniform
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Could not read file")
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var dictionary []string
	for scanner.Scan() {
    t := scanner.Text()
		if len(t) < 4 {
			// instructions: only 4 letters+ words matter
			continue
		}
		dictionary = append(dictionary, t)
	}

	var matches []string
	rows := strings.Split(wordSearch, "\n")
	for _, row := range rows {
		for _, word := range dictionary {
			if strings.Contains(row, word) {
				matches = append(matches, word)
			}
		}
	}

	var cols []string
	for i := 0; i < len(rows[0]); i++ {
		var sb strings.Builder
		for _, row := range rows {
			sb.WriteByte(row[i])
		}
		cols = append(cols, sb.String())
	}
	for _, col := range cols {
		for _, word := range dictionary {
			if strings.Contains(col, word) {
				matches = append(matches, word)
			}
		}
	}

  // remove duplicates
  seen := make(map[string]int, 0)
  var deduped []string
  for _, m := range matches {
    if seen[m] == 0 {
      deduped = append(deduped, m)
    }
    seen[m] = 1
  }

	fmt.Println(deduped)
  fmt.Println(len(deduped))
}

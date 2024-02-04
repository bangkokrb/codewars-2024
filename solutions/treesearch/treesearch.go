package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Id       int    `json:"id"`
	Children []Node `json:"children"`
}

type Search struct {
	History []Node
	Current *Node
}

// id -> sum
type Results map[int]int

func mapSlice[T any, M any](enum []T, mapper func(T) M) []M {
	n := make([]M, len(enum))
	for i, e := range enum {
		n[i] = mapper(e)
	}
	return n
}

func visit(history []Node, current Node, results Results) {
	// this is an endpoint
	if len(current.Children) == 0 {
		ids := mapSlice[Node, int](history, func(node Node) int {
			return node.Id
		})
		ids = append(ids, current.Id)

		sum := 0
		for _, n := range ids {
			sum += n
		}

		results[current.Id] = sum
	} else {
		for _, child := range current.Children {
			visit(append(history, current), child, results)
		}
	}
}

func main() {
	var root Node
	json.Unmarshal([]byte(jsChunk), &root)

	results := make(Results)
	visit([]Node{}, root, results)

	// fmt.Println(results)
	highid := 0
	highsum := 0

	for id, sum := range results {
		if sum > highsum {
			highid = id
			highsum = sum
		}
	}

	fmt.Printf("Highest ID: %d / sum: %d\n", highid, highsum)
	// Highest ID: 11947 / sum: 38506
}

const jsChunk = `
{"id":0,"children":[{"id":2263,"children":[{"id":11251,"children":[{"id":9626,
"children":[{"id":3010,"children":[{"id":760,"children":[]},{"id":8429,
"children":[]}]},{"id":422,"children":[{"id":9290,"children":[]},{"id":8230,
"children":[]},{"id":10213,"children":[]},{"id":8176,"children":[]}]},
{"id":3300,"children":[{"id":3689,"children":[]},{"id":10239,"children":[]},
{"id":11335,"children":[]},{"id":9722,"children":[]}]}]},{"id":9803,
"children":[{"id":1856,"children":[{"id":11593,"children":[]},{"id":3504,
"children":[]},{"id":1465,"children":[]},{"id":3328,"children":[]}]},
{"id":3759,"children":[{"id":8856,"children":[]},{"id":9201,"children":[]},
{"id":3797,"children":[]},{"id":10754,"children":[]}]}]},{"id":1448,
"children":[{"id":9812,"children":[{"id":8310,"children":[]},{"id":1816,
"children":[]},{"id":9665,"children":[]}]},{"id":11597,"children":[{"id":11947,
"children":[]},{"id":9397,"children":[]},{"id":8770,"children":[]}]},
{"id":1999,"children":[{"id":9772,"children":[]},{"id":1981,"children":[]},
{"id":11841,"children":[]},{"id":10164,"children":[]},{"id":12000,
"children":[]}]},{"id":913,"children":[]},{"id":3274,"children":[]}]},
{"id":976,"children":[{"id":1991,"children":[]},{"id":2052,"children":[]},
{"id":9184,"children":[]}]}]},{"id":9144,"children":[{"id":702,
"children":[{"id":8149,"children":[]},{"id":3765,"children":[]},{"id":8325,
"children":[]}]},{"id":284,"children":[{"id":2590,"children":[]},{"id":1990,
"children":[]}]},{"id":10370,"children":[{"id":10314,"children":[]},{"id":591,
"children":[]},{"id":10344,"children":[]},{"id":9682,"children":[]}]}]},
{"id":9718,"children":[{"id":2719,"children":[{"id":8546,"children":[]},
{"id":3261,"children":[]},{"id":9284,"children":[]}]},{"id":10775,
"children":[{"id":11488,"children":[]},{"id":10346,"children":[]},{"id":8132,
"children":[]},{"id":11316,"children":[]},{"id":10567,"children":[]}]}]}]},
{"id":1696,"children":[{"id":8884,"children":[{"id":1737,
"children":[{"id":8464,"children":[]},{"id":2723,"children":[]},{"id":9425,
"children":[]},{"id":2500,"children":[]}]},{"id":9195,"children":[{"id":618,
"children":[]},{"id":278,"children":[]},{"id":9048,"children":[]}]},{"id":273,
"children":[{"id":8043,"children":[]},{"id":36,"children":[]},{"id":2738,
"children":[]},{"id":10312,"children":[]},{"id":49,"children":[]}]},
{"id":11457,"children":[{"id":3768,"children":[]},{"id":1184,"children":[]},
{"id":2718,"children":[]}]},{"id":9464,"children":[{"id":11798,"children":[]},
{"id":835,"children":[]},{"id":1284,"children":[]},{"id":11066,"children":[]},
{"id":1203,"children":[]}]}]},{"id":8072,"children":[{"id":2928,
"children":[{"id":101,"children":[]},{"id":2230,"children":[]},{"id":9954,
"children":[]},{"id":2210,"children":[]}]},{"id":9156,"children":[{"id":3424,
"children":[]},{"id":10112,"children":[]},{"id":271,"children":[]},{"id":2944,
"children":[]}]}]},{"id":1621,"children":[{"id":3277,"children":[{"id":2563,
"children":[]},{"id":9215,"children":[]},{"id":10389,"children":[]},{"id":117,
"children":[]},{"id":1788,"children":[]}]},{"id":3903,"children":[{"id":1331,
"children":[]},{"id":3671,"children":[]},{"id":9411,"children":[]},{"id":9710,
"children":[]}]},{"id":8475,"children":[{"id":11203,"children":[]},{"id":11487,
"children":[]},{"id":1313,"children":[]},{"id":8569,"children":[]}]},{"id":100,
"children":[{"id":10193,"children":[]},{"id":8329,"children":[]}]}]},
{"id":9940,"children":[{"id":9897,"children":[{"id":3326,"children":[]},
{"id":2557,"children":[]},{"id":11426,"children":[]},{"id":444,
"children":[]}]},{"id":9230,"children":[{"id":46,"children":[]},{"id":2522,
"children":[]},{"id":8057,"children":[]},{"id":8644,"children":[]}]},
{"id":10202,"children":[{"id":354,"children":[]},{"id":10409,"children":[]},
{"id":982,"children":[]},{"id":1173,"children":[]},{"id":1040,
"children":[]}]}]},{"id":1005,"children":[{"id":3864,"children":[{"id":638,
"children":[]},{"id":8890,"children":[]},{"id":8978,"children":[]},{"id":9588,
"children":[]}]},{"id":10469,"children":[{"id":2000,"children":[]},{"id":8374,
"children":[]},{"id":11260,"children":[]},{"id":1310,"children":[]}]},
{"id":2609,"children":[{"id":1435,"children":[]},{"id":10827,"children":[]},
{"id":1636,"children":[]},{"id":9145,"children":[]}]},{"id":810,
"children":[{"id":9157,"children":[]},{"id":2934,"children":[]},{"id":3700,
"children":[]},{"id":10989,"children":[]}]},{"id":1792,"children":[{"id":10067,
"children":[]},{"id":1230,"children":[]}]}]}]},{"id":2029,
"children":[{"id":8049,"children":[{"id":11584,"children":[{"id":2261,
"children":[]},{"id":11815,"children":[]},{"id":1937,"children":[]},{"id":2328,
"children":[]}]},{"id":11776,"children":[{"id":9866,"children":[]},{"id":3353,
"children":[]},{"id":1183,"children":[]}]}]},{"id":11732,
"children":[{"id":11614,"children":[{"id":2108,"children":[]},{"id":8935,
"children":[]},{"id":11868,"children":[]}]},{"id":426,"children":[{"id":8406,
"children":[]},{"id":2382,"children":[]},{"id":11645,"children":[]},{"id":9246,
"children":[]}]},{"id":1722,"children":[{"id":10162,"children":[]},{"id":768,
"children":[]},{"id":10907,"children":[]}]},{"id":8190,"children":[{"id":1838,
"children":[]},{"id":2875,"children":[]},{"id":2834,"children":[]}]}]}]}]}
`

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("strategy.txt")
	if err != nil {
		log.Fatal(err)
	}

	doc := string(content)
	res := strings.Split(doc, "\n")
	var total int

	wins := make(map[string]string)
	wins["A"] = "C"
	wins["B"] = "A"
	wins["C"] = "B"

	fails := make(map[string]string)
	fails["C"] = "A"
	fails["A"] = "B"
	fails["B"] = "C"

	draw := make(map[string]string)
	draw["X"] = "A"
	draw["Y"] = "B"
	draw["Z"] = "C"

	values := make(map[string]int)
	values["A"] = 1
	values["B"] = 2
	values["C"] = 3

	for i := range res {
		result := string(res[i][2])
		their := string(res[i][0])

		// 	if wins[result]== their {
		// 		total += 6
		// 	} else if draw[result] == their {
		// 		total += 3
		// 	}
		// 	total += values[result]
		// }

		if result == "Y" {
			total += 3
			total += values[their]
		} else if result == "X" {
			my := wins[their]
			total += values[my]
		} else if result == "Z" {
			total += 6
			my := fails[their]
			total += values[my]
		}
	}
	fmt.Println(total)
}

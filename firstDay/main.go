package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"log"
	"os"
)

func main() {
//var maxSum int 
var SumForElf int
var maxThree [3]int
var sum int	

content, err := os.ReadFile("elves.txt")
if err != nil {
	log.Fatal(err)
}

doc := string(content)
res := strings.Split(doc, `

`)

for _, elf := range res {
		listOfSnacks := strings.Split(elf, "\n")
		for _, snack := range listOfSnacks {
			snackNum, _ := strconv.Atoi(snack)
			SumForElf += snackNum
		}
		// if SumForElf > maxSum {
		// 	maxSum = SumForElf
		// }
		sort.Ints(maxThree [:])
		if SumForElf > maxThree[0] {
			maxThree[0] = SumForElf
		}
		SumForElf = 0
	}

for _, v := range maxThree {
	sum += v
}

fmt.Println(sum)
}
	
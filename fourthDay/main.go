package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("sections.txt")
	if err != nil {
		log.Fatal(err)
	}

	doc := string(content)
	allSections := strings.Split(doc, "\n")

	var sumOverlap int

	for _, section := range allSections {
		twoElfs := strings.Split(section, ",")
		firstElf := twoElfs[0]
		secondElf := twoElfs[1]
		firstElfRange := strings.Split(firstElf, "-")
		secondElfRange := strings.Split(secondElf, "-")

		firstElfFirstValue, _ := strconv.Atoi(firstElfRange[0])
		firstElfSecondValue, _ := strconv.Atoi(firstElfRange[1])

		secondElfFirstValue, _ := strconv.Atoi(secondElfRange[0])
		secondElfSecondValue, _ := strconv.Atoi(secondElfRange[1])
		
		if !(firstElfSecondValue < secondElfFirstValue || secondElfSecondValue < firstElfFirstValue) {
			sumOverlap ++
		}
	}
	fmt.Println(sumOverlap)
	
}
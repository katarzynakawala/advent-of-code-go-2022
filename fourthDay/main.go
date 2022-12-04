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

	var sum int
	

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

		var firstElfFullRange []int
		
		for i := firstElfFirstValue; i <= firstElfSecondValue; i ++ {
			firstElfFullRange = append(firstElfFullRange, i)
		}

		var secondElfFullRange []int
		
		for i := secondElfFirstValue; i <= secondElfSecondValue; i ++ {
			secondElfFullRange = append(secondElfFullRange, i)
		}

		doesItContain := contains(firstElfFullRange, secondElfFullRange)

		if doesItContain {
			sum++
		}
	}
	fmt.Println(sum)
	
}

func contains(s1 []int, s2 []int) bool {

	var shorter []int
	var longer []int
	var count int 

	if len(s1) < len(s2) {
		shorter = s1
		longer = s2
		}else {
		shorter = s2
		longer = s1
		}
	
    for _, a := range shorter {
		for _, b := range longer {
			if a == b {
				count += 1
			}
		}
    } 
	if count == len(shorter){
		return true
	} else {
		return false
	}
}

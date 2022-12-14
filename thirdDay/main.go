package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("rucksacs.txt")
	if err != nil {
		log.Fatal(err)
	}

	doc := string(content)
	allRucksacks := strings.Split(doc, "\n")
	fmt.Println(allRucksacks[0])

	letter := map[rune]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 10, 'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15, 'p': 16, 'q': 17, 'r': 18, 's': 19, 't': 20, 'u': 21, 'v': 22, 'w': 23, 'x': 24, 'y': 25, 'z': 26, 'A': 27, 'B': 28, 'C': 29, 'D': 30, 'E': 31, 'F': 32, 'G': 33, 'H': 34, 'I': 35, 'J': 36, 'K': 37, 'L': 38, 'M': 39, 'N': 40, 'O': 41, 'P': 42, 'Q': 43, 'R': 44, 'S': 45, 'T': 46, 'U': 47, 'V': 48, 'W': 49, 'X': 50, 'Y': 51, 'Z': 52}
	var res int

	count := 0
	var oneRucksack string
	var twoRucksack string

	for _, rucksack := range allRucksacks {
		if count == 0 {
			oneRucksack = rucksack
			count += 1
		} else if count == 1 {
			twoRucksack = rucksack
			count += 1
		} else {
			threeRucksack := rucksack
			commonPrt := Intersection(oneRucksack, twoRucksack, threeRucksack)
			res += letter[commonPrt]
			count = 0
		}
	}
	fmt.Println(res)
}

func Intersection(a, b, c string) (e rune) {
	m := make(map[rune]bool)
	commonOneTwo := make(map[rune]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			commonOneTwo[item] = true
		}
	}
	for _, item := range c {
		if _, ok := commonOneTwo[item]; ok {
			e = item
		}
	}
	return
}

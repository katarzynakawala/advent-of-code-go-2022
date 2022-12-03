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


	letter := map[rune]int{'a':1, 'b':2, 'c':3, 'd':4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j':10, 'k':11, 'l':12, 'm':13, 'n':14, 'o':15, 'p':16, 'q':17, 'r':18, 's':19, 't':20, 'u':21, 'v':22, 'w':23, 'x':24, 'y':25, 'z':26, 'A':27, 'B':28, 'C':29, 'D':30, 'E': 31, 'F': 32, 'G': 33, 'H': 34, 'I': 35, 'J':36, 'K':37, 'L':38, 'M':39, 'N':40, 'O':41, 'P':42, 'Q':43, 'R':44, 'S':45, 'T':46, 'U':47, 'V':48, 'W':49, 'X':50, 'Y':51, 'Z':52 }
	var res int 

	for _, rucksack := range allRucksacks {
		halfRucksackLength := len(rucksack)/2
		firstCompartment := rucksack[:halfRucksackLength]
		secondCompartment := rucksack[halfRucksackLength:]

		commonPrt := Intersection(firstCompartment, secondCompartment)
		fmt.Println(string(commonPrt))
		
		withoutCopies := removeDuplicateRune(commonPrt)

		for _, l := range withoutCopies {
			res += letter[l]
		}
	}	
}

func Intersection(a, b string) (c []rune) {
	m := make(map[rune]bool) 

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
				c = append(c, item)
			}
		}
	return
}

func removeDuplicateRune(runeSlice []rune) []rune {
    allKeys := make(map[rune]bool)
    list := []rune{}
    for _, item := range runeSlice {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}
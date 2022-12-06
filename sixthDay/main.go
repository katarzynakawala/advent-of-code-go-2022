package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	doc := string(content)
	//marker := doc[:4]
	marker := doc[:14]

	for i:= len(marker); i < len(doc); i++ {
		if UnicChar(marker) {
			fmt.Println("Print ", i)
			break
		} else {
			fmt.Println("marker[1:]", marker[1:])
			fmt.Println("doc[i:i +1]", doc[i:i +1])
			marker = marker[1:] + doc[i:i +1]
			fmt.Println("Marker", marker)
		}
	}
}

func UnicChar(s string) bool {
	for i:= 0; i < len(s)- 1; i++ {
		for j:= i+ 1; j < len(s); j++ {
			if s[i:i+1] == s[j:j+1] {
				return false
			} 
		}
	}
	return true
}
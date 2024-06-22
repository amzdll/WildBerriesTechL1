package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	words := strings.Split(s, " ")
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	example := "snow dog cat sun"
	fmt.Println("source —", example)
	fmt.Print("reversed — ", reverse(example))
}

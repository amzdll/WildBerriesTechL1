package main

import (
	"fmt"
)

func toSet(s []string) map[string]struct{} {
	m := make(map[string]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}

func main() {
	example := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(toSet(example))
}

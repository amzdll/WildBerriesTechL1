package main

import (
	"fmt"
	"strings"
)

func checkUnique(s string) bool {
	m := map[int]struct{}{}
	for _, i := range strings.ToLower(s) {
		m[int(i)] = struct{}{}
	}
	return len(m) == len(s)
}

func main() {
	ExampleStrings := []string{"abcd", "abCdefAaf", "aabcd"}

	for i := range ExampleStrings {
		fmt.Println(ExampleStrings[i], "—", checkUnique(ExampleStrings[i]))
	}
}

package main

import "fmt"

func reverse(s string) string {
	convertedString := []rune(s)
	var letter rune
	for i := 0; i < len(convertedString)/2; i++ {
		letter = convertedString[i]
		convertedString[i] = convertedString[len(convertedString)-i-1]
		convertedString[len(convertedString)-i-1] = letter
	}
	return string(convertedString)
}

func main() {
	example := "главрыба"
	fmt.Println(example, "—", reverse(example))
}

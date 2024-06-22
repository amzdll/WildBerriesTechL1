package main

import "fmt"

func erase(s []interface{}, pos int) []interface{} {
	return append(s[:pos], s[pos+1:]...)
}

func main() {
	numsExample := []interface{}{1, 2, 3, 4, 5, 6}
	charExample := []interface{}{'1', '2', '3', '4', '5', '6'}
	stringExample := []interface{}{"111", "222", "333", "444", "555", "666"}
	fmt.Println(erase(numsExample, 1))
	fmt.Println(erase(charExample, 1))
	fmt.Println(erase(stringExample, 1))
}

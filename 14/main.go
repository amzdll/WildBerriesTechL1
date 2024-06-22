package main

import (
	"fmt"
	"os"
	"reflect"
)

func getOption() int {
	var option int
	fmt.Println("1 - integer")
	fmt.Println("2 - string")
	fmt.Println("3 - bool")
	fmt.Println("4 - channel")
	fmt.Fscan(os.Stdin, &option)
	return option
}

func defineType(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func main() {
	var intV int
	var stringV string
	var boolV bool
	var chanV chan int

	switch getOption() {
	case 1:
		fmt.Print(defineType(intV))
	case 2:
		fmt.Print(defineType(stringV))
	case 3:
		fmt.Print(defineType(boolV))
	case 4:
		fmt.Print(defineType(chanV))
	default:
		fmt.Print("Invalid option")
	}
}

package main

import (
	"errors"
	"fmt"
	"strconv"
)

func setBit(num *int, pos int, bit int) error {
	if pos < 0 || pos >= 64 {
		return errors.New("pos parameter must be in the range of 0 to 64")
	}
	mask := 1 << pos
	switch bit {
	case 0:
		*num &= ^mask
	case 1:
		*num |= mask
	default:
		return errors.New("invalid parameter (bit)")
	}
	return nil
}

func main() {
	example := 1
	pos := 0
	bit := 0

	fmt.Println("Before setting the bit - ", strconv.FormatInt(int64(example), 2))
	err := setBit(&example, pos, bit)
	fmt.Println("After setting the bit - ", strconv.FormatInt(int64(example), 2))
	if err != nil {
		fmt.Println(err.Error())
	}
}

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func binarySearch[T cmp.Ordered](s []T, target T) int {
	var mid int
	l, r := 0, len(s)-1

	for l <= r {
		mid = l + (r-l)/2
		if s[mid] < target {
			l = mid + 1
		} else if s[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	var stringExample = []string{"aa", "abc", "abcd"}
	slices.Sort(stringExample)
	fmt.Println(binarySearch(stringExample, "abcd"))

	var intExample = []int{1, 2, 3, 4}
	slices.Sort(intExample)
	fmt.Println(binarySearch(intExample, 4))
}

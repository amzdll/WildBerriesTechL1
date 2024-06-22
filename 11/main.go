package main

import "fmt"

func findCommon(m1, m2 map[interface{}]struct{}) map[interface{}]struct{} {
	res := make(map[interface{}]struct{})
	if len(m1) < len(m2) {
		for i := range m1 {
			if _, ok := m2[i]; ok {
				res[i] = struct{}{}
			}
		}
	} else {
		for i := range m2 {
			if _, ok := m1[i]; ok {
				res[i] = struct{}{}
			}
		}
	}
	return res
}

func main() {
	m1 := map[interface{}]struct{}{
		1: {}, 2: {}, 3: {}, 4: {},
	}
	m2 := map[interface{}]struct{}{
		3: {}, 4: {}, 5: {}, 6: {},
	}
	fmt.Println(findCommon(m1, m2))
}

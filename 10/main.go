package main

import (
	"fmt"
)

func groupTemperatures(temperatures []float64) map[int][]float64 {
	res := make(map[int][]float64)
	for _, temperature := range temperatures {
		key := int(temperature/10.0) * 10
		res[key] = append(res[key], temperature)
	}
	return res
}

func main() {
	temperatures := []float64{
		-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5,
	}
	fmt.Println(groupTemperatures(temperatures))
}

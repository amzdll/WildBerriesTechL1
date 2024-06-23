package main

import (
	"fmt"
	"math"
	"math/big"
)

func mulBig(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func divideBig(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func sumBig(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func subBig(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func main() {
	bigNum := int64(math.Pow(2.0, 21.0))
	a, b := big.NewInt(bigNum), big.NewInt(bigNum)

	fmt.Println(mulBig(a, b))
	fmt.Println(divideBig(a, b))
	fmt.Println(sumBig(a, b))
	fmt.Println(subBig(a, b))
}

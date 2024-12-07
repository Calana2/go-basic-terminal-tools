package main

import (
	"fmt"
	"main/student"
	"math/big"
)

func main() {
    tests := []struct {
        a, b *big.Int
    }{
        {big.NewInt(56), big.NewInt(15)},
        {big.NewInt(120), big.NewInt(23)},
        {big.NewInt(101), big.NewInt(10)},
        {big.NewInt(625), big.NewInt(90)},
        {big.NewInt(1001), big.NewInt(77)},
        {big.NewInt(270), big.NewInt(192)},
    }

    for _, test := range tests {
        mcd, x, y := student.Extended_Euclides(test.a, test.b)
        fmt.Printf("MCD(%s, %s) = %s\n", test.a.String(), test.b.String(), mcd.String())
        fmt.Printf("Coeficientes: x = %s, y = %s\n", x.String(), y.String())
    }
}


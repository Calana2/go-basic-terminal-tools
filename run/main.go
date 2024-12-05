package main

import (
	"fmt"
	"main/lcrypto"
	"math/big"
)

 func main() {
 // print(lcrypto.GCD(525,231))
 bigun := big.NewInt(2342343253463454)
 factors := lcrypto.Primefact(bigun)
 fmt.Printf("Prime factors of %v: %v\n",bigun,factors)

 // Example with 1:
 one := big.NewInt(1)
 factorsOne := lcrypto.Primefact(one)
 fmt.Println("Prime factors of 1:", factorsOne)

 // Example with 0:
    zero := big.NewInt(0)
    factorsZero := lcrypto.Primefact(zero)
    fmt.Println("Prime factors of 0:", factorsZero)
}

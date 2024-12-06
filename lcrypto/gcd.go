package lcrypto

import "math/big"

func GCD(a *big.Int, b *big.Int) *big.Int {
 // Euclides algorithm
 for b.Cmp(big.NewInt(0)) != 0 {
   tmp := new(big.Int).Set(b)
   b = new(big.Int).Mod(a,b)
   a = tmp
 }
 return a
}


package student

import "math/big"

func GCD(a *big.Int, b *big.Int) *big.Int {
 for b.Cmp(big.NewInt(0)) != 0 {
   tmp := new(big.Int).Set(b)
   b = new(big.Int).Mod(a,b)
   a = tmp
 }
 return a
}

func Extended_Euclides(a *big.Int, b *big.Int) (*big.Int,*big.Int,*big.Int) {
 x0, x1, y0, y1 := big.NewInt(1), big.NewInt(0), big.NewInt(0), big.NewInt(1)
 zero := big.NewInt(0)
 for b.Cmp(zero) != 0 {
   q := new(big.Int).Set(a).Div(a,b)
   tmp := new(big.Int).Set(b)
   b = new(big.Int).Mod(a,b)
   a = tmp
   x0, x1 = x1, new(big.Int).Sub(x0, new(big.Int).Mul(q, x1)) 
   y0, y1 = y1, new(big.Int).Sub(y0, new(big.Int).Mul(q, y1))
 }
 return a, x0, y0
}


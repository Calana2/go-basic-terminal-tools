package student

import (
	"math/big"
)

func Primefact(n *big.Int) (factors []*big.Int) {
// Division by test
  // case 0
  if n.Cmp(big.NewInt(1)) <= 0 {
   return
  }
  // case 2
  for n.Cmp(big.NewInt(1)) > 0 && new(big.Int).Rem(n, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
   factors = append(factors, big.NewInt(2))
   n.Div(n, big.NewInt(2))
  }
  // default
  limit := new(big.Int).Sqrt(n)
  divisor := big.NewInt(2)
  for divisor.Cmp(limit) <= 0 {
    for new(big.Int).Set(n).Rem(n,divisor).Cmp(big.NewInt(0)) == 0 {
      factors = append(factors, new(big.Int).Set(divisor)) 
      n.Div(n,divisor)
    }
   divisor.Add(divisor,big.NewInt(2))
  }
  // prime
  if n.Cmp(big.NewInt(1)) > 1 {
   factors = append(factors, new(big.Int).Set(n))
  }
  return factors
}

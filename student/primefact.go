package student

import (
	"math/big"
)

func Primefact(n *big.Int) (factors []*big.Int) {
  // Division by test
  if n.Cmp(big.NewInt(1)) <= 0 {
   return
  }

  divisor := big.NewInt(2)
  for n.Cmp(big.NewInt(1)) > 0 {
    for new(big.Int).Set(n).Rem(n,divisor).Cmp(big.NewInt(0)) == 0 {
      factors = append(factors, new(big.Int).Set(divisor)) 
      n.Div(n,divisor)
    }
   divisor.Add(divisor,big.NewInt(1))
  }
  return factors
}

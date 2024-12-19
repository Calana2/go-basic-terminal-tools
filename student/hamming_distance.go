package student

import "fmt"

func Hamming_distance(a []byte, b []byte) (int, error) {
 // Bit difference measure
 if len(a) != len(b) {
  return 0, fmt.Errorf("Params must have the same length ")
 }
 distance := 0
 for i:=0; i<len(a); i++ {
  xor := a[i] ^ b[i]
  for xor != 0 {
   distance += int(xor & 1)
   xor >>= 1
  }
 }
 return distance,nil
}

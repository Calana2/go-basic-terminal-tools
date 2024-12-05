package lcrypto

func GCD(a uint64, b uint64) uint64 {
 // Euclides algorithm
 for b != 0 {
   a, b = b, a % b
 }
 return a
}



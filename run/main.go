package main

import (
	"bufio"
	"io"
	"main/lcrypto"
  //"math/big"
  "os"
)

 func main() {
 // fmt.Print(lcrypto.GCD(big.NewInt(546234),big.NewInt(4352)))

 /*
  bigun := big.NewInt(2342343253463454)       // takes a while
  factors := lcrypto.Primefact(bigun)
  fmt.Printf("Prime factors of %v: %v\n",bigun,factors) 
 */

// /*
 var reader io.Reader

 if len(os.Args) > 2 {
  os.Exit(1)
 } else if len(os.Args) == 2 {
   file,err := os.Open(os.Args[1]) 
   defer file.Close()
   reader = file
  if err != nil {
   println(err)
   os.Exit(1)
  }
 } else {
  reader = os.Stdin
 }

  // pipe 
  if reader == os.Stdin {
   lcrypto.Base64Decode(*bufio.NewReader(reader))
   os.Exit(0)
  }
  // Normal input
  lcrypto.Base64Decode(*bufio.NewReader(reader)); os.Exit(0) 
}

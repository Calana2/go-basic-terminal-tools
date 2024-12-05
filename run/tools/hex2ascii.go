package main

import (
	"encoding/hex"
	"flag"
	"io"
	"os"
)

func main() {
 var reader io.Reader

 decode := flag.Bool("decode",false,"hex -> ASCII")
 flag.BoolVar(decode,"d",false,"")
 flag.Parse()


 if len(flag.Args()) > 0 {
  filePath :=  flag.Args()[0]
  file,err := os.Open(filePath)
  if err != nil {
   os.Exit(1)
  } 
  reader = file
  defer file.Close()
 } else {
  reader = os.Stdin
 }

 if *decode {
   decoder := hex.NewDecoder(reader)
   io.Copy(os.Stdout,decoder)
 } else {
   encoder := hex.NewEncoder(os.Stdout)
   io.Copy(encoder,reader)
 }
}

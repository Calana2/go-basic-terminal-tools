package main

import (
	b64 "encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
 var reader io.Reader

 decode := flag.Bool("decode",false,"base64 -> plain text")
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
  decoder := b64.NewDecoder(b64.StdEncoding,reader)
  _, err := io.Copy(os.Stdout,decoder)
  if err != nil {
   fmt.Println(err)
   os.Exit(1)
  } 
 } else {
   encoder := b64.NewEncoder(b64.StdEncoding,os.Stdout)
   _,err := io.Copy(encoder,reader)
   if err != nil {
    fmt.Println(err)
    os.Exit(1)
   }
   encoder.Close()
 }
}

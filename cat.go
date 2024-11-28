package main
import (
	"fmt"
	"io"
	"os"
)

func main() {

 var reader io.Reader

 if len(os.Args) > 1 {
  filePath :=  os.Args[1]
  file,err := os.Open(filePath)
  if err != nil {
   os.Exit(1)
  } 
  reader = file
  defer file.Close()
 } else {
  reader = os.Stdin
 }
 
 buffer := make([]byte,1024)
 for {
   n,err := reader.Read(buffer) 
   if err == io.EOF { break }
   if err != nil {
    fmt.Println(err)
    continue
   }
   fmt.Print(string(buffer[:n]))
 }
}

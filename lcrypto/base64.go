package lcrypto

import (
	"bufio"
	"fmt"
	"io"
)

const BASE64_MAP = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"


/* ----------------------------
  **********UTILITIES**********
  _____________________________ */

func bytes2bits (bytes []byte) []bool {
  var bits []bool
  for _, b := range bytes {
   for i:=7; i>=0; i-- {
     // bit masking
     bit := (b & (1 << i)) != 0
     bits = append(bits,bit)
   }
  }
  return bits
}

func bits2b64indices (bits []bool) []int {
 var indices []int
 for i:=0; i<len(bits); i+=6 {
   var value int
   for j:=0; j<6; j++ {
    if i+j < len(bits) && bits[i+j] {
     value |= 1 << (5-j)
    }
   }
   indices = append(indices,value)
 }
 return indices
}

func bytes2indices(bytes []byte) ([]int,error) {
 indexMap := make(map[byte]int)
 for i,b := range BASE64_MAP {
  indexMap[byte(b)] = i
 }

 indices := make([]int,len(bytes))
 for n,b := range bytes {
  index, ok := indexMap[b]
  if !ok {
    return nil, fmt.Errorf("Byte %d does not found in the map table",b)
  }
  indices[n] = index
 }
 return indices,nil 
}

func indices2bytes (indices []int) (bytes [3]byte) {
 bits := (indices[0] << 18) | (indices[1] << 12) | (indices[2] << 6) | indices[3]  
 bytes[0] = byte((bits >> 16) & 0xff)
 bytes[1] = byte((bits >> 8) & 0xff)
 bytes[2] = byte(bits & 0xff)
 return bytes
}

func RemoveNewlines(chunk []byte) (newChunk []byte, count int) { 
  for _, b := range chunk { 
    if b != 0xA { 
      newChunk = append(newChunk, b) 
    } else {
      count++
    }
  }
  for len(newChunk) < len(chunk) { 
    newChunk = append(newChunk, 0x0) 
  }
  return newChunk, count
}

/* --------------------------
  **********ENCODER**********
  ___________________________ */

func Base64Encode(reader io.Reader) {
// Encode accepts a io.Reader and converts it to base64, if you want to pass them a string you must convert it
// Ex: Encode(strings.NewReader("mystring"))
 chunk := make([]byte,3)
 padding := 0
 char_count := 0 
 line := ""
 for {
  n,err := reader.Read(chunk)  
  if err == io.EOF { break }
  if err != nil { continue }
  var bit_chunk []bool
  // Add padding
  if n < 3 {
   padding = 3 - n
   for i:= n; i<3 ;i++ {
    chunk[i]=0
   }
  }
   // Convert 3 bytes to 24 bits
   bit_chunk = bytes2bits(chunk) 
  // Convert bits to integer
  indices := bits2b64indices(bit_chunk)
  // Print 
  for i,val := range(indices) {
  // Print the padding bytes
   if padding != 0 && val == 0 && i >= 4 - padding {
    line+="="
  // Print the mapped bytes
   } else {
    line+=string(BASE64_MAP[val]) 
   }
   char_count++
  }
  if char_count == 76 { fmt.Println(line); char_count=0; line="" }
 }
 if len(line) > 0 {
  fmt.Println(line)
 }
}



/* ----------------------------
  **********DECODER**********
  _____________________________ */

func Base64Decode(reader bufio.Reader) {
 // Decode accepts a io.Reader, if you want to pass them a string you must convert it
 // Ex: Decode(strings.NewReader("Q29kZW1vbmtleXM=")) 
 chunk := make([]byte,4)
 char_count := 0 
 line := ""
 for {
  _,err := reader.Read(chunk)  
  if err == io.EOF { break }
  if err != nil { continue } 
  // Remove new line chars
  chunk,count := RemoveNewlines(chunk)
  for i:=0; i<count; i++ {
   b,_ := reader.ReadByte()
   chunk[3-i] = b 
  }
  // Convert 4 bytes to map indices
  indices,err := bytes2indices(chunk)
  if err != nil { fmt.Println(err); break }
  // Convert indices to 24 bits and bits to bytes
  bytes := indices2bytes(indices)
  line += string(bytes[:])
  if char_count == 76 { fmt.Print(line); char_count=0; line="" }
 }
 if len(line) > 0 {
  fmt.Print(line)
 }
}

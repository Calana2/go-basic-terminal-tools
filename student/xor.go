package student
import "fmt"


func Fixed_xor(var1 []byte, var2 []byte) ([]byte, error) {
  var result []byte
  if len(var1) != len(var2) {
    return nil, fmt.Errorf("Params must have the same length")
  }
  for i:=0; i < len(var1); i++ {
    result[i] = var1[i] ^ var2[i]
  }
  return result,nil
}

func Repeated_xor(msg []byte, key []byte) ([]byte, error) {
  var result []byte
  if len(msg) < len(key) {
    return nil, fmt.Errorf("Key must be less or equal than the message")
  }
  key_index := 0
  for i:=0; i < len(msg); i++ {
    result[i] = msg[i] ^ key[key_index]
    key_index = (key_index+1) % len(key)
  }
  return result,nil
}

func Single_xor(msg []byte, key byte) []byte {
  var result []byte
  for i:=0; i < len(msg); i++ {
    result[i] = msg[i] ^ key
  }
  return result
}







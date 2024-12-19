package student

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
)

func AES_ECB_Encrypt(key, data []byte) (string, error) {
  // to base64
  ciph, err := aes.NewCipher(key)
  if err != nil {
    return "",err
  }
  blockSize := ciph.BlockSize()
  p_data := pkcs7Padding(data,blockSize)
  ciphertext := make([]byte,len(p_data))
  for i, j := 0, 16; i < len(p_data); i, j = i+16,j+16 {
    ciph.Encrypt(ciphertext[i:j],p_data[i:j])
  }
  return base64.RawStdEncoding.EncodeToString(ciphertext),nil
}

func pkcs7Padding(data []byte, blockSize int) []byte {
  padding := blockSize - len(data) % blockSize
  padtext := bytes.Repeat([]byte{byte(padding)},padding)
  return append(data,padtext...)
}

func AES_ECB_Decrypt(key []byte, b64_data string) (string, error) {
  // from base64
  ciph, err := aes.NewCipher(key)
  if err != nil {
    return "",err
  }
  p_data,err := base64.RawStdEncoding.DecodeString(b64_data)
  if err != nil {
    return "",err
  }
  plaintext := make([]byte,len(p_data))
  for i, j := 0, 16; i < len(p_data); i, j = i+16,j+16 {
    ciph.Decrypt(plaintext[i:j],p_data[i:j])
  }
  return string(pkcs7Unpadding(plaintext)),nil
}

func pkcs7Unpadding(data []byte) []byte {
  padding := int(data[len(data)-1])
  return data[:len(data)-padding]
}

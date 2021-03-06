package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	plainText := []byte("Bob loves Alice.")
	key := []byte("passw0rdpassw0rdpassw0rdpassw0rd")

	// Create new DES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}

	// The IV (Initialization Vector) need to be unique, but not secure.
	// Therefore, it's common to include it at the beginning of the cipher text.
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	// Create IV
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Printf("err: %s\n", err)
	}

	// Encrypt
	encryptStream := cipher.NewCFBEncrypter(block, iv)
	encryptStream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	fmt.Printf("Cipher text: %v\n", cipherText)

	// Decrpt
	decryptedText := make([]byte, len(cipherText[aes.BlockSize:]))
	decryptStream := cipher.NewCFBDecrypter(block, cipherText[:aes.BlockSize])
	decryptStream.XORKeyStream(decryptedText, cipherText[aes.BlockSize:])
	fmt.Printf("Decrypted text: %s\n", string(decryptedText))
}

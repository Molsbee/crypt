package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

func main() {
	// Encrypt File
	data, _ := os.ReadFile("temp.txt")

	block, _ := aes.NewCipher([]byte("TestingSomething"))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	os.WriteFile("ciphertext.bin", ciphertext, 0777)

	// Decrypt Encrypted File
	ciphertext, _ = os.ReadFile("ciphertext.bin")
	nonce = ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]
	plaintext, _ := gcm.Open(nil, nonce, ciphertext, nil)
	os.WriteFile("plaintext.txt", plaintext, 0777)
}

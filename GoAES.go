package main

import (
	//"crypto/aes"
	//"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// 128 bit encryption method & cipher
func main() {
	fmt.Println("Beginning GoAES program")
	initial_text := "Hello"
	hashed_text := shaHash(initial_text)

	fmt.Println("Initial: ", initial_text)
	fmt.Println("Hashed (sha256): ", hashed_text)
	fmt.Println("Encrypted: ")
	fmt.Println("Decrypted: ")
}

func shaHash(input string) string {
	plain_text_input := []byte(input)
	hashed_input := sha256.Sum256(plain_text_input)
	return hex.EncodeToString(hashed_input[:])
}

func encryptMe(value []byte, key_phrase string) []byte {
	return
}

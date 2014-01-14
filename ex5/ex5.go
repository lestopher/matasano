package main

import (
	"encoding/hex"
	"fmt"
)

func encrypt(s, key []byte) string {
	cipher := make([]byte, len(s))

	for i := range s {
		cipher[i] = s[i] ^ key[i%len(key)]
	}

	return hex.EncodeToString(cipher)
}

func main() {
	s := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := []byte("ICE")

	fmt.Println(encrypt(s, key))
}

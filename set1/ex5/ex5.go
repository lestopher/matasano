package ex5

import "encoding/hex"

func Encrypt(s, key []byte) string {
	cipher := make([]byte, len(s))

	for i := range s {
		cipher[i] = s[i] ^ key[i%len(key)]
	}

	return hex.EncodeToString(cipher)
}

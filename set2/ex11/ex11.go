package ex11

import "crypto/rand"

func RandomKey(byteSize int) []byte {
	key := make([]byte, byteSize)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return key
}

func EncryptionOracle(plaintext []byte) []byte {

}

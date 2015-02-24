package ex10

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"

	"github.com/lestopher/matasano/set2/ex9"
)

func EncryptCBC(key, message, iv []byte) []byte {
	block, _ := aes.NewCipher(key)

	// TODO: pad the message
	if len(message) < aes.BlockSize {
		fmt.Println("message is smaller than blocksize")
		message = []byte(ex9.Pad(message, aes.BlockSize))
	}

	if len(iv) != aes.BlockSize {
		panic("iv is not the same length as aes block size")
	}

	if len(message)%aes.BlockSize != 0 {
		fmt.Println("message is not a multiple of the block size")
		message = []byte(ex9.Pad(message, aes.BlockSize))
	}

	cryptedMsg := make([]byte, len(message))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cryptedMsg, message)

	return cryptedMsg
}

func DecryptCBC(key, message, iv []byte) []byte {
	block, _ := aes.NewCipher(key)

	// TODO: pad the message
	if len(message) < aes.BlockSize {
		fmt.Println("message is smaller than blocksize")
		message = []byte(ex9.Pad(message, aes.BlockSize))
	}

	if len(iv) != aes.BlockSize {
		fmt.Printf("expected length %d, got %d\n", aes.BlockSize, len(iv))
		// panic("iv is not the same length as aes block size")
		iv = iv[:aes.BlockSize]
	}

	if len(message)%aes.BlockSize != 0 {
		fmt.Println("message is not a multiple of the block size")
		message = []byte(ex9.Pad(message, aes.BlockSize))
	}

	decryptedMsg := make([]byte, len(message))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decryptedMsg, message)

	return decryptedMsg
}

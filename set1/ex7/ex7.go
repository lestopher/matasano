package ex7

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"

	"github.com/lestopher/matasano/set2/ex9"
)

func DecryptECB(key, message []byte) []byte {
	block, _ := aes.NewCipher(key)

	if len(message) < aes.BlockSize {
		fmt.Println("message is smaller than blocksize")
		message = []byte(ex9.Pad(message, aes.BlockSize))
	}

	iv := make([]byte, aes.BlockSize)

	if len(message)%aes.BlockSize != 0 {
		fmt.Println("message is not a multiple of the block size, padding")
		message = []byte(ex9.Pad(message, aes.BlockSize))
	}

	decryptedMsg := make([]byte, aes.BlockSize)
	decryptedBlock := make([]byte, aes.BlockSize)

	// Is there a better way to do this if there's no initialization vector?
	for i := 0; i < len(message); i += aes.BlockSize {
		mode := cipher.NewCBCDecrypter(block, iv)
		mode.CryptBlocks(decryptedBlock, message[i:i+aes.BlockSize])
		decryptedMsg = append(decryptedMsg, decryptedBlock...)
	}

	return decryptedMsg[aes.BlockSize:]
}

func EncryptECB(key, message []byte) []byte {
	block, _ := aes.NewCipher(key)

	if len(message) < aes.BlockSize {
		panic("message is smaller than blocksize")
	}

	iv := make([]byte, aes.BlockSize)

	if len(message)%aes.BlockSize != 0 {
		panic("message is not a multiple of the block size")
	}

	encryptedMsg := make([]byte, aes.BlockSize)
	encryptedBlock := make([]byte, aes.BlockSize)

	for i := 0; i < len(message); i += aes.BlockSize {
		mode := cipher.NewCBCEncrypter(block, iv)
		mode.CryptBlocks(encryptedBlock, message[i:i+aes.BlockSize])
		encryptedMsg = append(encryptedMsg, encryptedBlock...)
	}
	fmt.Printf("%v\n", encryptedMsg[:aes.BlockSize*2])
	return encryptedMsg[aes.BlockSize*2:]
}

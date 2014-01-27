package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	var input string
	key := []byte("YELLOW SUBMARINE")
	file, err := os.Open("../data/ex7_gistfile.txt")
	defer file.Close()

	if err != nil {
		panic("Couldn't open file.")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input += scanner.Text()
	}

	byteArray, b64Err := base64.StdEncoding.DecodeString(input)

	if b64Err != nil {
		fmt.Errorf("Error decoding base64 string to byteArray: %s\n", b64Err)
	}

	block, _ := aes.NewCipher(key)

	if len(byteArray) < aes.BlockSize {
		panic("byteArray is smaller than blocksize")
	}

	iv := make([]byte, aes.BlockSize)
	// iv := byteArray[:aes.BlockSize]
	// byteArray = byteArray[aes.BlockSize:]

	if len(byteArray)%aes.BlockSize != 0 {
		panic("byteArray is not a multiple of the block size")
	}

	// mode := cipher.NewCBCDecrypter(block, iv)

	answer := make([]byte, aes.BlockSize)
	temp := make([]byte, aes.BlockSize)

	for i := 0; i < len(byteArray); i += aes.BlockSize {
		mode := cipher.NewCBCDecrypter(block, iv)
		mode.CryptBlocks(temp, byteArray[i:i+aes.BlockSize])

		answer = append(answer, temp...)
	}

	fmt.Printf("%s\n", answer)
}

package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	. "fmt"
	"os"
)

func unique(u []byte) ([]byte, int) {
	const BLOCKSIZE = 16
	counter := make(map[string]int)
	temp := make([]byte, BLOCKSIZE)
	var uniqueSlice []byte

	for i := 0; i < len(u); i += BLOCKSIZE {
		temp = u[i : i+BLOCKSIZE]
		_, exists := counter[string(temp)]

		if exists {
			counter[string(temp)]++
		} else {
			counter[string(temp)] = 1
		}
	}

	for key, value := range counter {
		if value == 1 {
			uniqueSlice = append(uniqueSlice, []byte(key)...)
		}
	}

	return uniqueSlice, len(uniqueSlice)
}

func main() {
	file, _ := os.Open("../data/ex8_gistfile.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ba [][]byte

	for scanner.Scan() {
		raw, _ := hex.DecodeString(scanner.Text())
		ba = append(ba, raw)
	}

	for i, b := range ba {
		lenB := len(b)
		_, l := unique(b)

		if lenB != l {
			Println("Has dupes:", i)
			Printf("Expected size: %d, got %d\n", lenB, l)
		}
	}

	key := []byte("YELLOW SUBMARINE")
	block, _ := aes.NewCipher(key)
	iv := make([]byte, aes.BlockSize)

	firstPart := ba[132][:aes.BlockSize]
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(firstPart, firstPart)

	Println(string(firstPart))
}

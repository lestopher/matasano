package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/lestopher/matasano/set1/ex3"
	"github.com/lestopher/matasano/set1/ex6"
	"os"
)

func main() {
	file, err := os.Open("./data/ex6_gistfile.txt")
	var input, likelyKey string

	if err != nil {
		panic("Couldn't open file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}

	if err != nil {
		fmt.Println("Error: ", err)
	}

	byteArray, b64Err := base64.StdEncoding.DecodeString(input)

	if b64Err != nil {
		fmt.Errorf("Error decoding base64 string to byteArray: %s\n", b64Err)
	}

	probableKeysizes := ex6.Top5Keysizes(byteArray)
	probableByteArrays := make([]ex6.BlockCollection, 6)

	for i, ks := range probableKeysizes {
		blockCollection, _ := ex6.ToBlockCollection(byteArray, ks.Keysize)
		transposedBlockCollection, _ := ex6.TransposeBlocks(blockCollection)
		probableByteArrays[i] = transposedBlockCollection
	}

	for _, tbc := range probableByteArrays {
		for _, block := range tbc {
			s := ex3.Cleanup(ex3.BestGuess([]byte(block)))

			if len(s) > 0 {
				likelyKey += string(s[0].Key)
			}
		}
		if len(likelyKey) > 0 {
			break
		}
	}

	fmt.Println("Key is", likelyKey)
}

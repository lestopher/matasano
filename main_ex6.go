package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"matasano/ds"
	"matasano/ex6"
	"os"
)

func main() {
	file, err := os.Open("./data/gistfile2.txt")
	var input, likelyKey string
	// var input string

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

	// Returns the most probable keysize and the hamming distance
	key, _ := ex6.SmallestHamdist(byteArray)

	fmt.Println("key is", key)
	blockCollection, _ := ex6.ToBlockCollection(byteArray, 4)

	transposedBlockCollection, _ := ex6.TransposeBlocks(blockCollection)

	for _, block := range transposedBlockCollection {
		// s := ds.BestGuess(block)
		likelyKey += string(ds.BestGuess(block)[0].Key)
	}

	fmt.Println("Key is", likelyKey)
	fmt.Println(string(ds.Decrypt(byteArray, likelyKey)))
}

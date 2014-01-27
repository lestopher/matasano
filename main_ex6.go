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
	ksc := ex6.Top5Keysizes(byteArray)
	fmt.Println(ksc)

	fmt.Println("key is", key)
	blockCollection, bcLength := ex6.ToBlockCollection(byteArray, 29)
	fmt.Println("bcLength is", bcLength)

	transposedBlockCollection, tbcLength := ex6.TransposeBlocks(blockCollection)
	fmt.Println("tbcLength is", tbcLength)

	for i, block := range transposedBlockCollection {
		fmt.Println("length block", len(block))
		s := ds.Cleanup(ds.BestGuess(block))
		if i == 24 {
			fmt.Println(s[0].Dstring)
		}
		// s := ds.BestGuess(block)
		// fmt.Println(ds.Cleanup(s))
		if len(s) > 0 {
			likelyKey += string(s[0].Key)
		} else {
			likelyKey += "^"
		}
		// likelyKey += string(s[0].Key)
	}

	fmt.Println("Key is", likelyKey)
	// fmt.Println(string(ds.Decrypt(byteArray, likelyKey)))
}

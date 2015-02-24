// A consumer file for example 2
// XOR two strings and prints the results
package main

import (
	"encoding/hex"
	"fmt"

	"github.com/lestopher/matasano/set1/ex2"
)

func main() {
	const stringOne = "1c0111001f010100061a024b53535009181c"
	const stringTwo = "686974207468652062756c6c277320657965"

	hexStringOne, _ := hex.DecodeString(stringOne)
	hexStringTwo, _ := hex.DecodeString(stringTwo)

	fmt.Println(hex.EncodeToString(ex2.Xor(hexStringOne, hexStringTwo)))
}

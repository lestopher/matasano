package main

import (
	"encoding/hex"
	"fmt"
	"matasano/ex3"
)

func main() {
	var dsc ex3.DecipheredStringCollection
	cipher, err := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	if err != nil {
		panic("Whoopsie, something went super wrong.")
	}

	dsc = ex3.BestGuess(cipher)
	fmt.Printf("Key: %s\nDeciphered String - %s\nChi Square Score: %f\n", dsc[0].Key, dsc[0].Dstring, dsc[0].ChiSquareScore)
}

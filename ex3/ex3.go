// package ex3
package main

import (
	"encoding/hex"
	"fmt"
	"matasano/ds"
)

func main() {
	var dsc ds.DecipheredStringCollection
	cipher, err := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	if err != nil {
		panic("Whoopsie, something went super wrong.")
	}

	dsc = ds.BestGuess(cipher)
	fmt.Println(dsc)
	fmt.Println(dsc[0].Dstring)
}

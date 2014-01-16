// package ex3
package main

import (
	"encoding/hex"
	"fmt"
	"matasano/ds"
	"sort"
)

func main() {
	var dsc ds.DecipheredStringCollection
	cipher, err := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	if err != nil {
		panic("Whoopsie, something went super wrong.")
	}

	for _, v := range ds.ALPHA_ARRAY {
		s := string(ds.Decrypt(cipher, v))
		d := ds.NewDecipheredString(s, v, ds.ChiSquareSum(s))
		dsc = append(dsc, *d)
	}

	sort.Sort(sort.Reverse(dsc))

	fmt.Printf("KEY - %s - %s\n", dsc[0].Key, dsc[0].String)
}

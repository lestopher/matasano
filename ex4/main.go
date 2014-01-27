package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"matasano/ds"
	"os"
)

func main() {
	var inputDataArr [][]byte
	var dsc ds.DecipheredStringCollection

	file, err := os.Open("../data/gistfile.txt")

	if err != nil {
		panic("Couldn't open file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input, _ := hex.DecodeString(scanner.Text())
		inputDataArr = append(inputDataArr, input)
	}

	dsc = ds.BestGuessOnCollection(inputDataArr)
	// fmt.Printf("%f - %s - KEY - %s\n", dsc[0].ChiSquareScore, dsc[0].Key, dsc[0].String)
	fmt.Println(ds.Cleanup(dsc)[0].Dstring)
	// fmt.Println(dsc[:16])
}

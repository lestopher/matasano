package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"matasano/ex3"
	"os"
)

func main() {
	var inputDataArr [][]byte
	var dsc ex3.DecipheredStringCollection

	file, err := os.Open("./data/ex4_gistfile.txt")

	if err != nil {
		panic("Couldn't open file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input, _ := hex.DecodeString(scanner.Text())
		inputDataArr = append(inputDataArr, input)
	}

	dsc = ex3.BestGuessOnCollection(inputDataArr)
	fmt.Println(ex3.Cleanup(dsc)[0].Dstring)
}

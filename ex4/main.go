package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"matasano/ds"
	"os"
	"sort"
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

	for _, byteArray := range inputDataArr {
		for i := 0; i < 128; i++ {
			s := string(ds.Decrypt(byteArray, string(i)))
			d := ds.NewDecipheredString(s, string(i), ds.ChiSquareSum(s))
			dsc = append(dsc, *d)
		}
	}

	sort.Sort(sort.Reverse(dsc))
	dsc = ds.Cleanup(dsc)
	fmt.Printf("%f - %s - KEY - %s\n", dsc[0].ChiSquareScore, dsc[0].Key, dsc[0].String)
}

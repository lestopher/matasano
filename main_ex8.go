package main

import (
	"bufio"
	"encoding/hex"
	. "fmt"
	"github.com/lestopher/matasano/set1/ex8"
	"os"
)

func main() {
	file, _ := os.Open("./data/ex8_gistfile.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ba [][]byte

	for scanner.Scan() {
		raw, _ := hex.DecodeString(scanner.Text())
		ba = append(ba, raw)
	}

	// We're checking the unique length vs the real length, if it's different
	// we most likely have a AES EBC
	for i, b := range ba {
		lenB := len(b)
		uniq := ex8.Unique(b)

		if lenB != len(uniq) {
			Println("Has dupes:", i)
			Printf("Expected size: %d, got %d\n", lenB, len(uniq))
		}
	}
}

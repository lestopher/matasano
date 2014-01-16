package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../data/gistfile2.txt")
	var input string

	if err != nil {
		panic("Couldn't open file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input += scanner.Text()
	}

	if err != nil {
		fmt.Println("Error: ", err)
	}

	byteArray, b64Err := base64.StdEncoding.DecodeString(input)

	if b64Err != nil {
		fmt.Errorf("Error decoding base64 string to byteArray: %s\n", b64Err)
	}

}

// A consumer file for example 1
package main

import (
	. "fmt"
	"matasano/ex1"
)

func main() {
	const testString = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	Println(ex1.HexToBase64(testString))
}

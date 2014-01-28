package main

import (
	. "fmt"
	"matasano/ex5"
)

func main() {
	const testString = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	const key = "ICE"

	Println(ex5.Encrypt([]byte(testString), []byte(key)))
}

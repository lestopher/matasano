package main

import (
	. "fmt"
	"github.com/lestopher/matasano/set1/ex5"
)

func main() {
	const testString = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	const key = "ICE"

	Println(ex5.Encrypt([]byte(testString), []byte(key)))
}

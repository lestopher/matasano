package ex6

import (
	// "bytes"
	// "encoding/binary"
	// "fmt"
	// "math/big"
	// "encoding/base64"
	// "encoding/hex"
	// "matasano/ds"
	// "matasano/ex1"
	"matasano/ex2"
	// "matasano/ex5"
)

func Hamdist(x, y []byte) int {
	dist := 0
	// An xor gives us a result that tells us which bits were different
	val := ex2.Xor(x, y)

	for _, val_uint := range val {
		// From wikipedia, it sums all the active bits per byte slice
		for val_uint > 0 {
			dist++
			val_uint &= val_uint - 1
		}
	}

	return dist
}

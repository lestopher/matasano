package ex6

import (
	"fmt"
	"matasano/ex2"
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

func SmallestHamdist(s []byte) (int, float64) {
	var res float64
	var potentialK int
	var potentialN float64

	for KEYSIZE := 2; KEYSIZE <= 128; KEYSIZE++ {
		// if len(s)%KEYSIZE == 0 {
		fmt.Println(KEYSIZE)
		bytes := [][]byte{s[0:KEYSIZE], s[KEYSIZE : KEYSIZE*2], s[KEYSIZE*2 : KEYSIZE*3], s[KEYSIZE*3 : KEYSIZE*4]}

		for outerIndex, outerVal := range bytes {
			for _, innerVal := range bytes[outerIndex:] {
				res += NormalizedHamdist(Hamdist(outerVal, innerVal), KEYSIZE)
			}
		}

		res /= 6.

		if res < potentialN || KEYSIZE == 2 {
			potentialK = KEYSIZE
			potentialN = res
		}
		// }
	}

	return potentialK, potentialN
}

func NormalizedHamdist(hamdist, keysize int) float64 {
	return float64(hamdist) / float64(keysize)
}

func TransposeBlocks(blocks [][]byte) ([][]byte, int) {
	// A transposed block's length is the same as the length of the first element
	// you're transposing
	transposedBlocks := make([][]byte, len(blocks[0]))

	for i := range transposedBlocks {
		transposedBlocks[i] = make([]byte, len(blocks))
		for j := range blocks {
			transposedBlocks[i][j] = blocks[j][i]
		}
	}

	return transposedBlocks, len(transposedBlocks)
}

func ToBlockCollection(byteArray []byte, keysize int) ([][]byte, int) {
	if len(byteArray)%keysize != 0 {
		panic("byteArray is not evenly divisible by keysize")
	}

	collection := make([][]byte, len(byteArray)/keysize)
	collectionIndex := 0

	for i := 0; i < len(byteArray); i += keysize {
		collection[collectionIndex] = byteArray[i : i+keysize]
		collectionIndex++
	}

	return collection, len(collection)
}

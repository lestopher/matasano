package ex6

import (
	"fmt"
	"github.com/lestopher/matasano/set1/ex2"
	"sort"
)

type BlockCollection [][]byte

type Keysize struct {
	Keysize  int
	NHamdist float64
}

type KeysizeCollection []Keysize

func (k KeysizeCollection) Len() int {
	return len(k)
}

func (k KeysizeCollection) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func (k KeysizeCollection) Less(i, j int) bool {
	return k[i].NHamdist < k[j].NHamdist
}

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
	}

	return potentialK, potentialN
}

func Top5Keysizes(s []byte) KeysizeCollection {
	var res float64
	var ksc KeysizeCollection

	for KEYSIZE := 2; KEYSIZE <= 40; KEYSIZE++ {
		bytes := [][]byte{s[0:KEYSIZE], s[KEYSIZE : KEYSIZE*2], s[KEYSIZE*2 : KEYSIZE*3], s[KEYSIZE*3 : KEYSIZE*4]}

		for outerIndex, outerVal := range bytes {
			for _, innerVal := range bytes[outerIndex:] {
				res += NormalizedHamdist(Hamdist(outerVal, innerVal), KEYSIZE)
			}
		}

		res /= 6.

		ksc = append(ksc, Keysize{KEYSIZE, res})
	}

	sort.Sort(ksc)

	return ksc[0:6]
}

func NormalizedHamdist(hamdist, keysize int) float64 {
	return float64(hamdist) / float64(keysize)
}

func TransposeBlocks(blocks BlockCollection) (BlockCollection, int) {
	// For an MxN matrix, you need to create an NxM matrix
	// Then loop through and fill in the data
	mLen := len(blocks)
	nLen := len(blocks[0])

	// The "N" part of it
	transposedBlocks := make(BlockCollection, nLen)
	tbContent := make([]byte, nLen*mLen)

	for i := range transposedBlocks {
		transposedBlocks[i], tbContent = tbContent[:mLen], tbContent[mLen:]
	}

	for i := range blocks {
		for j := range transposedBlocks {
			transposedBlocks[j][i] = blocks[i][j]
		}
	}

	return transposedBlocks, len(transposedBlocks)
}

func ToBlockCollection(byteArray []byte, keysize int) (BlockCollection, int) {
	ba := byteArray

	if len(ba)%keysize != 0 {
		fmt.Println("byteArray length not divisible by keysize, zerofilling.")
		zerofill := make([]byte, keysize-(len(ba)%keysize))
		ba = append(ba, zerofill...)
	}

	collection := make(BlockCollection, len(ba)/keysize)
	collectionIndex := 0

	for i := 0; i < len(ba); i += keysize {
		collection[collectionIndex] = ba[i : i+keysize]
		collectionIndex++
	}

	return collection, len(collection)
}

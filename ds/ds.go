package ds

import (
	"math"
	"strings"
)

const alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789,./<>?;':\"[]\\{}|`~!@#$%^&*()_+-="

var ALPHA_ARRAY []string = strings.Split(alphabet, "")

// Frequency distributions of english letters a-z
var frequencies = []float64{
	.0817, .0149, .0278, .0425, .1270, .0223, .0202, .0609, .0697, .0015, .0077, .0402, .0241,
	.0675, .0751, .0193, .0009, .0599, .0633, .0906, .0276, .0098, .0236, .0015, .0197, .0007,
}

type DecipheredString struct {
	String, Key    string
	ChiSquareScore float64
}

func NewDecipheredString(deciphered_string string, key string, chi_square_score float64) (ds *DecipheredString) {
	ds = &DecipheredString{
		String:         deciphered_string,
		Key:            key,
		ChiSquareScore: chi_square_score,
	}

	return ds
}

type DecipheredStringCollection []DecipheredString

// Implements Sort interface
func (dsc DecipheredStringCollection) Len() int {
	return len(dsc)
}

func (dsc DecipheredStringCollection) Less(one, two int) bool {
	return dsc[one].ChiSquareScore > dsc[two].ChiSquareScore
}

func (dsc DecipheredStringCollection) Swap(one, two int) {
	dsc[one], dsc[two] = dsc[two], dsc[one]
}

func Cleanup(dsc DecipheredStringCollection) DecipheredStringCollection {
	var good DecipheredStringCollection
	var itsBad bool = false
	for _, ds := range dsc {
		for _, letter := range ds.String {
			if letter < 0x20 || letter > 0x7e {
				itsBad = true
			}
			if letter < 0x20 && (letter == 0x0a || letter == 0x0d || letter == 0x09) {
				itsBad = false
			}
		}
		if !itsBad {
			if float64(LengthCharsOnly(ds.String))/float64(len(ds.String)) > 0.74 {
				good = append(good, ds)
			}
		}
		itsBad = false
	}
	return good
}

func Decrypt(cipher []byte, key string) []byte {
	result := make([]byte, len(cipher))
	for i, c := range cipher {
		result[i] = c ^ []byte(key)[0]
	}
	return result
}

func ChiSquareSum(c string) float64 {
	letterMap := make(map[rune]int)
	var sum float64

	// Get a count of how many times a letter occurs
	for _, char := range c {
		char |= 0x20
		if 'a' <= char && char <= 'z' {
			_, ok := letterMap[char]
			if ok {
				letterMap[char] += 1
			} else {
				letterMap[char] = 1
			}
		}
	}

	for key, value := range letterMap {
		expectedfreq := frequencies[(int(key)-int('a'))%26] * float64(len(c))
		sum += math.Pow((float64(value)-expectedfreq), 2) / expectedfreq
	}

	return sum
}

func LengthCharsOnly(c string) int {
	length := 0
	for _, char := range c {
		char |= 0x20
		if 'a' <= char && char <= 'z' {
			length += 1
		}
	}
	return length
}

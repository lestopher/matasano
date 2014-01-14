package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"sort"
	"strings"
)

// const alphabet string = "abcdefghijklmnopqrstuvwxyz"

// const alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

const alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789,./<>?;':\"[]\\{}|`~!@#$%^&*()_+-="

var alpha_array []string = strings.Split(alphabet, "")

// Frequency distributions of english letters a-z
var frequencies = []float64{
	.0817, .0149, .0278, .0425, .1270, .0223, .0202, .0609, .0697, .0015, .0077, .0402, .0241,
	.0675, .0751, .0193, .0009, .0599, .0633, .0906, .0276, .0098, .0236, .0015, .0197, .0007,
}

type DecipheredString struct {
	string, key, original_string string
	chi_square_score             float64
}

type DecipheredStringCollection []DecipheredString

// Len, Less, Swap implements Sort interface
func (slice DecipheredStringCollection) Len() int {
	return len(slice)
}

func (slice DecipheredStringCollection) Less(one, two int) bool {
	return slice[one].chi_square_score > slice[two].chi_square_score
}

func (slice DecipheredStringCollection) Swap(one, two int) {
	slice[one], slice[two] = slice[two], slice[one]
}

func (slice DecipheredStringCollection) String() {
	for i := range slice {
		fmt.Printf("%f - %s - %s\n", slice[i].chi_square_score, slice[i].string, slice[i].key)
	}
}

func decrypt(cipher []byte, key string) []byte {
	result := make([]byte, len(cipher))
	for i, c := range cipher {
		// c |= 0x20
		result[i] = c ^ []byte(key)[0]
	}
	return result
}

func len_chars_only(c string) int {
	length := 0
	for _, char := range c {
		char |= 0x20
		if 'a' <= char && char <= 'z' {
			length += 1
		}
	}
	return length
}

func chi_square_sum(c string) float64 {
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

func main() {
	var dsc DecipheredStringCollection
	var inputByteArray [][]byte
	// var bestGuess DecipheredStringCollection
	var input string

	for {
		_, err := fmt.Scanln(&input)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		is, _ := hex.DecodeString(input)
		inputByteArray = append(inputByteArray, is)
	}

	for _, byteArray := range inputByteArray {
		for i := 0; i < 128; i++ {
			s := string(decrypt(byteArray, string(i)))
			ds := DecipheredString{string: s, key: string(i), chi_square_score: chi_square_sum(s), original_string: hex.EncodeToString(byteArray)}
			dsc = append(dsc, ds)
		}
	}
	sort.Sort(sort.Reverse(dsc))
	debugPrint(cleanup(dsc))
}

func debugPrint(c DecipheredStringCollection) {
	// For debugging
	for i := range c {
		fmt.Printf("%f - %s - KEY - %s - %s\n", c[i].chi_square_score, c[i].original_string, c[i].key, c[i].string)
	}
}

func cleanup(dsc DecipheredStringCollection) DecipheredStringCollection {
	var good DecipheredStringCollection
	var itsBad bool = false
	for _, ds := range dsc {
		for _, letter := range ds.string {
			if letter < 0x20 || letter > 0x7e {
				itsBad = true
			}
			if letter < 0x20 && (letter == 0x0a || letter == 0x0d || letter == 0x09) {
				itsBad = false
			}
		}
		if !itsBad {
			if float64(len_chars_only(ds.string))/float64(len(ds.string)) > 0.74 {
				good = append(good, ds)
			}
		}
		itsBad = false
	}
	return good
}

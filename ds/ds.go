package ds

import "strings"

const alphabet string = "abcdefghijklmnopqrstuvwxyz"

var alpha_array []string = strings.Split(alphabet, "")

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
func (slice DecipheredStringCollection) Len() int {
	return len(slice)
}

func (slice DecipheredStringCollection) Less(one, two int) bool {
	return slice[one].ChiSquareScore > slice[two].ChiSquareScore
}

func (slice DecipheredStringCollection) Swap(one, two int) {
	slice[one], slice[two] = slice[two], slice[one]
}

func Decrypt(cipher []byte, key string) []byte {
	result := make([]byte, len(cipher))
	for i := range cipher {
		result[i] = (cipher[i] ^ 0x20) ^ []byte(key)[0]
	}
	return result
}

func ChiSquareSum(c string) float64 {
	length := len_chars_only(c)
	letterMap := make(map[rune]int)
	var sum float64

	// Get a count of how many times a letter occurs
	for _, char := range c {
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
		expectedFreq := frequencies[(int(key)-int('a'))%26] * float64(length)
		sum += math.Pow((float64(value)-expectedFreq), 2) / expectedFreq
	}

	return sum
}

func LengthCharsOnly(c string) int {
	length := 0
	for _, char := range c {
		if 'a' <= char && char <= 'z' {
			length += 1
		}
	}
	return length
}

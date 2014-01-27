package ds

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

const alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789,./<>?;':\"[]\\{}|`~!@#$%^&*()_+-="

var ALPHA_ARRAY []string = strings.Split(alphabet, "")

// Frequency distributions of english letters a-z
var frequencies = []float64{
	0.08167, 0.01492, 0.02782, 0.04253, 0.12702, 0.02228, 0.02015, 0.06094, 0.06966, 0.00153, 0.00772, 0.04025, 0.02406,
	0.06749, 0.07507, 0.01929, 0.00095, 0.05987, 0.06327, 0.09056, 0.02758, 0.00978, 0.02360, 0.00150, 0.01974, 0.00074,
}

type DecipheredString struct {
	Dstring, Key   string
	ChiSquareScore float64
}

func (ds DecipheredString) String() string {
	return fmt.Sprintf("%s: %f", ds.Key, ds.ChiSquareScore)
}

func NewDecipheredString(deciphered_string string, key string, chi_square_score float64) (ds *DecipheredString) {
	ds = &DecipheredString{
		Dstring:        deciphered_string,
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
	return dsc[one].ChiSquareScore < dsc[two].ChiSquareScore
}

func (dsc DecipheredStringCollection) Swap(one, two int) {
	dsc[one], dsc[two] = dsc[two], dsc[one]
}

func Cleanup(dsc DecipheredStringCollection) DecipheredStringCollection {
	var good DecipheredStringCollection
	var itsBad bool = false
	for _, ds := range dsc {
		tildeCount := 0
		backtickCount := 0

		for _, letter := range ds.Dstring {
			if letter < 0x20 || letter > 0x7e {
				itsBad = true
			}

			// OK symbols \n, \r, \t
			if letter < 0x20 && (letter == 0x0a || letter == 0x0d || letter == 0x09) {
				itsBad = false
			}

			// tilde ~
			if letter == 0x7e {
				tildeCount++
				if tildeCount > 3 {
					itsBad = true
				}
			}

			// backtick `
			if letter == 0x60 {
				backtickCount++
			}
		}

		// If we have an uneven amount of backticks, probably bad
		if backtickCount%2 != 0 {
			itsBad = true
		}
		if !itsBad {
			// Check the ratio of letters to symbols... arbitrarily chose 60%
			if float64(LengthCharsOnly(ds.Dstring))/float64(len(ds.Dstring)) > 0.6 {
				good = append(good, ds)
			}
		}
		itsBad = false
	}
	return good
}

func Decrypt(cipher []byte, key string) []byte {
	result := make([]byte, len(cipher))
	keyByteArray := []byte(key)

	for i, c := range cipher {
		result[i] = c ^ keyByteArray[i%len(keyByteArray)]
	}
	return result
}

func ChiSquareSum(c string) (float64, float64) {
	var englishDist, uniformDist float64
	counts := make([]int, 26)
	lengthChars := float64(LengthCharsOnly(c))

	// Get a count of how many times a letter occurs
	for _, char := range c {
		char |= 0x20
		if 'a' <= char && char <= 'z' {
			counts[int(char)-97]++
		}
	}

	for i := 0; i < 26; i++ {
		englishDist += math.Pow((float64(counts[i])-(lengthChars*frequencies[i])), 2) / (lengthChars * frequencies[i])
		uniformDist += math.Pow((float64(counts[i])-(lengthChars/26.0)), 2) / (lengthChars / 26.0)
	}

	return englishDist, uniformDist
}

func LengthCharsOnly(c string) int {
	length := 0
	for _, char := range c {
		char |= 0x20
		if 'a' <= char && char <= 'z' {
			length++
		}
	}
	return length
}

func BestGuessOnCollection(input [][]byte) DecipheredStringCollection {
	var dsc DecipheredStringCollection

	for _, byteArray := range input {
		for i := 32; i < 128; i++ {
			s := string(Decrypt(byteArray, string(i)))
			cs, _ := ChiSquareSum(s)
			d := NewDecipheredString(s, string(i), cs)
			dsc = append(dsc, *d)
		}
		// dsc = Cleanup(dsc)
	}

	sort.Sort(dsc)

	return dsc
}

func BestGuess(input []byte) DecipheredStringCollection {
	var dsc DecipheredStringCollection

	for i := 32; i < 128; i++ {
		s := string(Decrypt(input, string(i)))
		cs, _ := ChiSquareSum(s)
		d := NewDecipheredString(s, string(i), cs)
		dsc = append(dsc, *d)
		// dsc = Cleanup(dsc)
	}

	sort.Sort(dsc)

	return dsc
}

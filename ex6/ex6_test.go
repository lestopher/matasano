package ex6

import "testing"
import "fmt"

func TestHamdist(t *testing.T) {
	var one = []byte{1, 0, 1, 1, 1, 0, 1}
	var two = []byte{1, 0, 0, 1, 0, 0, 1}
	const dist_one_two = 2

	var ms1 = []byte("this is a test")
	var ms2 = []byte("wokka wokka!!!")
	const dist_ms1_ms2 = 37

	result := Hamdist(one, two)

	if result != dist_one_two {
		t.Errorf("Expected weight of %d, got %d\n", dist_one_two, result)
	} else {
		t.Log("Byte array test passed.")
	}

	result = Hamdist(ms1, ms2)

	if result != dist_ms1_ms2 {
		t.Errorf("Expected weight of %d, got %d\n", dist_ms1_ms2, result)
	} else {
		t.Log("String comparison passed.")
	}
}

func TestToBlockCollection(t *testing.T) {
	const keysize int = 2
	b := []byte("hello world")
	blockCollection := ToBlockCollection(b, keysize)
	fmt.Println(blockCollection)
}

// func TestTransposedBlocks(t *testing.T) {
// 	matrix1 := [][]byte{[1, 2]}
// 	matrix1expected := [][]byte{[1], [2]}
// 	matrix2 := [][]byte{[1, 2], [3, 4]}
// 	matrix2expected := [][]byte{[1, 3], [2, 4]}
// 	matrix3 := [][]byte{[1, 2], [3, 4], [5, 6]}
// 	matrix3expected := [][]byte{[1, 3, 5], [2, 4, 6]}

// }
// func TestSmallestHamdist(t *testing.T) {
// G
// }

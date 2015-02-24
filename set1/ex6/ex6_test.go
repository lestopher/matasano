package ex6

import "testing"
import "fmt"

func TestHamdist(t *testing.T) {
	var one = []byte{1, 0, 1, 1, 1, 0, 1}
	var two = []byte{1, 0, 0, 1, 0, 0, 1}
	const DistOneTwo = 2

	var ms1 = []byte("this is a test")
	var ms2 = []byte("wokka wokka!!!")
	const DistMs1Ms2 = 37

	result := Hamdist(one, two)

	if result != DistOneTwo {
		t.Errorf("Expected weight of %d, got %d\n", DistOneTwo, result)
	} else {
		t.Log("Byte array test passed.")
	}

	result = Hamdist(ms1, ms2)

	if result != DistMs1Ms2 {
		t.Errorf("Expected weight of %d, got %d\n", DistMs1Ms2, result)
	} else {
		t.Log("String comparison passed.")
	}
}

func TestToBlockCollection(t *testing.T) {
	const keysize int = 2
	b := []byte("hello world")
	blockCollection, colLen := ToBlockCollection(b, keysize)
	if colLen%keysize != 0 {
		t.Errorf("Expected colLen(%d) mod keysize(%d) to be 0\n", colLen, keysize)
	}
	fmt.Println(blockCollection)
}

package ex2

import (
	"encoding/hex"
	"testing"
)

func TestXor(t *testing.T) {
	const stringOne = "1c0111001f010100061a024b53535009181c"
	const stringTwo = "686974207468652062756c6c277320657965"
	const expectedString = "746865206b696420646f6e277420706c6179"

	var result string

	hexStringOne, err1 := hex.DecodeString(stringOne)
	hexStringTwo, err2 := hex.DecodeString(stringTwo)

	if err1 != nil {
		panic(err1)
	}

	if err2 != nil {
		panic(err2)
	}

	result = hex.EncodeToString(Xor(hexStringOne, hexStringTwo))

	if result != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, result)
	}
}

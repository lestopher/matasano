package ex9

import (
	"bytes"
	"testing"
)

func TestPad(t *testing.T) {
	const blockSize = 20
	var correctlyPadded = []byte("YELLOW SUBMARINE\x04\x04\x04\x04")
	var msg = []byte("YELLOW SUBMARINE")
	padded := Pad(msg, blockSize)

	if len(padded) != blockSize {
		t.Errorf("Expected padded to be length %d, got %d\n", blockSize, len(padded))
	}

	if !bytes.Equal(padded, correctlyPadded) {
		t.Errorf("Expected %s received %s\n", correctlyPadded, padded)
	}
}

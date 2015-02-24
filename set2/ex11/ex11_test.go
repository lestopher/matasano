package ex11

import (
	"fmt"
	"testing"
)

func TestRandomKey(t *testing.T) {
	const keySize = 16
	key := RandomKey(keySize)

	fmt.Printf("Generated key:\t%x\n", key)

	if len(key) != keySize {
		t.Errorf("Expected key size to be: %d\nGot key size: %d\n", keySize, len(key))
	}
}

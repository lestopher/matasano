package ex1

import (
	"fmt"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	const testString = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	const expectedString = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	var base64String string

	base64String = HexToBase64(testString)

	if base64String != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, base64String)
	}
}

func TestBase64ToHex(t *testing.T) {
	const testKey = "FSlWccMpIZeY3aFeOV1f8A=="
	const testIv = "wzQyTn4wTs/RC8FYkwNUeA=="
	base64Key := Base64ToHex(testKey)
	base64Iv := Base64ToHex(testIv)
	fmt.Println(base64Key, len(base64Key))
	fmt.Println(base64Iv, len(base64Iv))
}

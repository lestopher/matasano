package ex1

import "testing"

func TestHexToBase64(t *testing.T) {
	const testString = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	const expectedString = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	var base64String string

	base64String = HexToBase64(testString)

	if base64String != expectedString {
		t.Errorf("Expected %s, got %s\n", expectedString, base64String)
	}
}

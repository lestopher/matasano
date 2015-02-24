package ex1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

/**
 * Converts a string representation of a hex value into a base64 encoded string
 * @param hs String The hex string
 * @return String base64 encoded string
 */
func HexToBase64(hs string) string {
	hexValue, err := hex.DecodeString(hs)

	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(hexValue)
}

func Base64ToHex(s string) string {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Println(s)
		panic(err)
	}
	return hex.EncodeToString(decoded)
}

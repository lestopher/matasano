package ex10

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)

func TestEncryptCBC(t *testing.T) {
	iv := []byte{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00',
		'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'}
	key := []byte("YELLOW SUBMARINE")
	cryptedMsg := readFileToBytes("/Users/christophernguyen/go/src/github.com/lestopher/matasano/data/10.txt")
	decryptedMsg := DecryptCBC(key, cryptedMsg, iv)
	reencryptedMsg := EncryptCBC(key, decryptedMsg, iv)

	if !bytes.Equal(cryptedMsg, reencryptedMsg) {
		t.Errorf("\nExpected\n-----\n%v\nGot\n-----\n%v\n", cryptedMsg, reencryptedMsg)
	}
}

func readFileToBytes(filePath string) []byte {
	var input string
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		panic("Couldn't open file.")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input += scanner.Text()
	}

	byteArray, b64Err := base64.StdEncoding.DecodeString(input)

	if b64Err != nil {
		fmt.Errorf("error decoding base64 string to byteArray: %s\n", b64Err)
	}

	return byteArray

}

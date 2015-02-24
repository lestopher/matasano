package ex7

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestDecryptECB(t *testing.T) {
	const path = "/Users/christophernguyen/go/src/github.com/lestopher/matasano/data/"
	const cryptedFile = "ex7_gistfile.txt"
	const expectedFile = "ex7_expected.txt"
	key := []byte("YELLOW SUBMARINE")
	msg := readBase64FileToBytes(path + cryptedFile)

	// Decrypt the message then recencrypt to compare to original msg
	decryptedMsg := DecryptECB(key, msg)
	reencryptedMsg := EncryptECB(key, decryptedMsg)

	if !bytes.Equal(msg, reencryptedMsg) {
		t.Errorf("\nExpected\n-----\n%v\nGot\n-----\n%v\n", msg, reencryptedMsg)
	}
}

func readFileToBytes(filePath string) []byte {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return b
}
func readBase64FileToBytes(filePath string) []byte {
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

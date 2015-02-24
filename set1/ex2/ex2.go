package ex2

/**
 * Takes two byte arrays and xors them and returns result
 * @param one []byte byte array one
 * @param two []byte byte array two
 * @return []byte
 */
func Xor(one, two []byte) []byte {
	if len(one) != len(two) {
		panic("length mismatch")
	}

	result := make([]byte, len(one))

	for i := range one {
		result[i] = one[i] ^ two[i]
	}

	return result
}

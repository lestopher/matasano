package ex9

// Pad will pad a byte array to the corresponding block size with char padding
func Pad(msg []byte, blockSize int) []byte {
	paddingSize := blockSize - (len(msg) % blockSize)
	result := make([]byte, len(msg)+paddingSize)
	copy(result, msg)
	for i := len(msg); i < len(result); i++ {
		result[i] = byte(paddingSize)
	}

	return result
}

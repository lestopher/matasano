package ex8

func Unique(u []byte) []byte {
	const BLOCKSIZE = 16
	counter := make(map[string]int)
	temp := make([]byte, BLOCKSIZE)
	var uniqueSlice []byte

	for i := 0; i < len(u); i += BLOCKSIZE {
		temp = u[i : i+BLOCKSIZE]
		_, exists := counter[string(temp)]

		if exists {
			counter[string(temp)]++
		} else {
			counter[string(temp)] = 1
		}
	}

	for key, value := range counter {
		if value == 1 {
			uniqueSlice = append(uniqueSlice, []byte(key)...)
		}
	}

	return uniqueSlice
}

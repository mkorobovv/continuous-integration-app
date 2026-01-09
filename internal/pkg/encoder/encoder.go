package encoder

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Encode(number int64) string {
	if number == 0 {
		return string(charset[0])
	}

	var chars []byte
	for number > 0 {
		chars = append(chars, charset[number%62])
		number /= 62
	}

	// Reverse the result
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}

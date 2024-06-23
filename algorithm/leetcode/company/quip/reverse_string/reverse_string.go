package reverse_string

import "strings"

func Reverse(s string, separator string) string {
	bytes := []byte(s)
	reverseBytes(bytes)
	prev := -1
	for i := 0; i < len(bytes); i++ {
		if strings.Contains(separator, string(bytes[i])) {
			reverseBytes(bytes[prev+1 : i])
			prev = i
		}
	}
	reverseBytes(bytes[prev+1:])
	return string(bytes)
}

func reverseBytes(bytes []byte) {
	i, j := 0, len(bytes)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
}

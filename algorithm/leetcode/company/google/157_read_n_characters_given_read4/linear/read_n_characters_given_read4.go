package linear

import "math"

/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */
var solution = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	return func(buf []byte, n int) int {
		// The return value (how many bytes we actually read)
		char := 0
		// How many bytes we read by calling read4. If less than 4, it means we hit the file end
		tmpChar := 4
		// the input buffer ptr
		i := 0
		// if we still want to read more bytes and the are bytes left
		for n > 0 && tmpChar == 4 {
			tmpBuf := make([]byte, 4)
			tmpChar = read4(tmpBuf)
			// actual bytes we want to process
			actualChar := minInt(tmpChar, n)
			char += actualChar
			for j := 0; j < actualChar; j++ {
				buf[i+j] = tmpBuf[j]
			}
			i += actualChar
			n -= actualChar
		}
		return char
	}
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

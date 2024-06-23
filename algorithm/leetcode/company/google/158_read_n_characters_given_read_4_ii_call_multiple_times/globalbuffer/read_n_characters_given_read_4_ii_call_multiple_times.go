package globalbuffer

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
	// buffer to store the results of read 4 because we probably are not able to consume
	// all bytes from read4. If not, we need to cache the remain bytes in the buffer.
	var buffer = make([]byte, 4)
	// current buffer pointer
	var bufferPtr = 0
	// # of bytes in buffer
	var bufferCnt = 0
	// implement read below.
	return func(buf []byte, n int) int {
		// return value
		numBytes := 0
		bufPtr := 0
		for n > 0 {
			// we always read the bytes into global buffer
			if bufferCnt == 0 {
				bufferCnt = read4(buffer)
			}
			// get the actual bytes we can/need process
			actualChar := minInt(bufferCnt-bufferPtr, n)
			numBytes += actualChar
			// if after read bufferCnt is 0, it mean we hit the file end and we need to break
			if bufferCnt == 0 {
				break
			}
			i := 0
			// copy
			for bufferPtr < bufferCnt && i < actualChar {
				buf[bufPtr+i] = buffer[bufferPtr]
				bufferPtr++
				i++
			}
			bufPtr += actualChar
			n -= actualChar
			// if bufferPtr hits the buffer end, we need to update bufferCnt and bufferPtr to be
			// 0 (refresh the buffer)
			if bufferPtr == bufferCnt {
				bufferCnt = 0
				bufferPtr = 0
			}
		}
		return numBytes
	}
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

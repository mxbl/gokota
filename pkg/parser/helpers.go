package parser

import "bytes"

// Find the last line in `buf` that entirely consists of `sep` and return index
func findLastLine(buf []byte, sep string) int {
	bs := []byte(sep)
	max := len(buf)
	for {
		i := bytes.LastIndex(buf[:max], bs)
		if i <= 0 || i+len(bs) >= len(buf) {
			return -1
		}
		if (buf[i-1] == '\n' || buf[i-1] == '\r') &&
			(buf[i+len(bs)] == '\n' || buf[i+len(bs)] == '\r') {
			return i
		}
		max = i
	}
}

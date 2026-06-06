package main

import "unsafe"

// intToString returns a string from an integer without allocating.
// The result is valid until the next call to intToString.
func intToString(i int) string {
	if i == 0 {
		return "0"
	}
	n := len(intBuf)
	for i > 0 {
		n--
		intBuf[n] = byte(i%10) + '0'
		i /= 10
	}
	return unsafe.String(unsafe.SliceData(intBuf[n:]), len(intBuf)-n)
}

// stringToInt returns an integer from a string without having to use strconv package.
func stringToInt(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		result = result*10 + (int(s[i]) - 48)
	}

	return result
}

// bytesToInt returns an integer from a byte slice without allocating.
func bytesToInt(b []byte) int {
	result := 0

	for i := 0; i < len(b); i++ {
		result = result*10 + (int(b[i]) - 48)
	}

	return result
}

// uintToBytes writes the decimal representation of v into dst and returns the number of bytes written.
func uintToBytes(dst []byte, v uint32) int {
	if v == 0 {
		dst[0] = '0'
		return 1
	}
	var tmp [10]byte
	n := 0
	for v > 0 {
		tmp[n] = byte(v%10) + '0'
		n++
		v /= 10
	}
	for i := 0; i < n; i++ {
		dst[i] = tmp[n-1-i]
	}
	return n
}

// Package cobs implenments COBS (https://en.wikipedia.org/wiki/Consistent_Overhead_Byte_Stuffing).
package cobs

// #include <stddef.h>
// #include <stdint.h>
// #include "cobs.h"
// #cgo CFLAGS: -g -Wall -flto
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

// Encode encodes `i` into `o` and returns number of bytes in `o`.
// For details see https://en.wikipedia.org/wiki/Consistent_Overhead_Byte_Stuffing.
// The Encode implementation is done in C because the aimed use case is an embedded device running C.
// This function is mainly for testing.
func Encode(o, i []byte) (n int) {
	if len(i) == 0 {
		return
	}
	in := unsafe.Pointer(&i[0])
	out := unsafe.Pointer(&o[0])
	n = int(C.COBSEncode(out, in, C.size_t(len(i))))
	return
}

// decodeCOBS expects in slice rd a byte sequence ending with a 0, writes the COBS decoded data to wr and returns len(wr).
//
// If rd contains more bytes after the first 0 byte, these are ignored.
// Needs to be written in a better way.
func decodeCOBS(wr, rd []byte) int {
	n, e := Decode(wr, rd)
	if e == nil {
		return n
	}
	fmt.Println(e)
	return 0
}

// Decode a COBS frame to a slice of bytes.
//
// decoded := dec[:n]
func Decode(dec, cobs []byte) (n int, e error) {

	for len(cobs) > 0 {
		cnt := cobs[0] - 1
		cobs = cobs[1:]
		if int(cnt) > len(cobs) {
			e = errors.New("inconsistent COBS packet")
			return
		}
		n += copy(dec[n:], cobs[:cnt])
		cobs = cobs[cnt:]
		if cnt < 0xfe && len(cobs) > 0 {
			dec[n] = 0
			n++
		}
	}
	return
}

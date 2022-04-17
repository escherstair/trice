package src

import (
	"fmt"
	"io"
	"testing"

	"github.com/tj/assert"
)

// TestTriceSequencesNoFraming ...
func TestTriceSequencesNoFraming(t *testing.T) {
	out := make([]byte, 1024)
	setTriceBuffer(out)
	setTriceFraming(0)

	for i, exp := range triceBytesNoFrame {
		len := triceCode(i)
		act := out[:len]
		//fmt.Printf("i=%d, ", i)
		//dump(os.Stdout, act)
		assert.Equal(t, exp, act)
	}
}

// triceBytesNoFrame contains the expected trice byte streams when C.TriceCode is executed.
// This is not encoded.
var triceBytesNoFrame = [][]byte{
	{0x22, 0x62, 0xc0, 0x4, 0xff, 0xff, 0xff, 0xff},
	{0x22, 0xa2, 0x11, 0x11, 0xc0, 0x4, 0xff, 0xff, 0xff, 0xff},
	{0x22, 0xe2, 0x11, 0x11, 0x11, 0x11, 0xc0, 0x4, 0xff, 0xff, 0xff, 0xff},
}

// TestTriceSequencesCOBSFraming ...
func TestTriceSequencesCOBSFraming(t *testing.T) {
	out := make([]byte, 1024)
	setTriceBuffer(out)
	setTriceFraming(1)

	for i, exp := range triceBytesCOBSFrame {
		len := triceCode(i)
		act := out[:len]
		//fmt.Printf("i=%d, ", i)
		//dump(os.Stdout, act)
		assert.Equal(t, exp, act)
	}
}

// triceBytesNoFrame contains the expected trice byte streams when C.TriceCode is executed.
// This is not encoded.
var triceBytesCOBSFrame = [][]byte{
	{9, 0x22, 0x62, 0xc0, 0x4, 0xff, 0xff, 0xff, 0xff, 0},
	{11, 0x22, 0xa2, 0x11, 0x11, 0xc0, 0x4, 0xff, 0xff, 0xff, 0xff, 0},
	{13, 0x22, 0xe2, 0x11, 0x11, 0x11, 0x11, 0xc0, 0x4, 0xff, 0xff, 0xff, 0xff, 0},
}

// TestTriceSequencesTCOBSFraming ...
func TestTriceSequencesTCOBSFraming(t *testing.T) {
	out := make([]byte, 1024)
	setTriceBuffer(out)
	setTriceFraming(2)

	for i, exp := range triceBytesTCOBSFrame {
		len := triceCode(i)
		act := out[:len]
		//fmt.Printf("i=%d, ", i)
		//dump(os.Stdout, act)
		assert.Equal(t, exp, act)
	}
}

// triceBytesTCOBSFrame contains the expected trice byte streams when C.TriceCode is executed.
// This is not encoded.
var triceBytesTCOBSFrame = [][]byte{
	{0x22, 0x62, 0xc0, 0x4, 0x84, 0},
	{0x22, 0xa2, 0x11, 0x11, 0xc0, 0x4, 0x86, 0},
	{0x22, 0xe2, 0x11, 0x13, 0xc0, 0x4, 0x82, 0},
}

// TRICE32_1( Id(58755), "rd:TRICE32_1 line %d (%%d)\n", -1 );
//
//  "58755": {
//  	"Type": "TRICE32_1",
//  	"Strg": "rd:TRICE32_1 line %d (%%d)\\n"
//  },
//
// []byte{0x02, 0x03, 0x01, 0x01, 0x02, 0x1a, 0x0f, 0x37, 0xcb, 0x11, 0x11, 0x11, 0x11, 0xc0, 0x01, 0x83, 0xe5, 0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00}
//
// "rd:TRICE32_1 line -1 (%d)\n"

// dump prints the byte slice as hex in one line
func dump(w io.Writer, b []byte) {
	fmt.Fprint(w, "exp := []byte{ ")
	for _, x := range b {
		fmt.Fprintf(w, "0x%02x, ", x)
	}
	fmt.Fprintln(w, "}")
}

//! \file cobs.c
//! \author Thomas.Hoehenleitner [at] seerose.net


#include <stdint.h>
#include <stddef.h>
#include "cobs.h"

//! COBSEncode encodes data to output.
//! @param in Pointer to input data to encode.
//! @param length Number of bytes to encode.
//! @param out Pointer to encoded output buffer.
//! @return Encoded buffer length in bytes.
//! @note Does not output delimiter byte. Code taken from Wikipedia and slightly adapted.
size_t COBSEncode( void * restrict out, const void * restrict in, size_t length) {
    uint8_t * buffer = out;
	uint8_t *encode = buffer; // Encoded byte pointer
	uint8_t *codep = encode++; // Output code pointer
	uint8_t code = 1; // Code value

	for (const uint8_t *byte = (const uint8_t *)in; length--; ++byte)
	{
		if (*byte) { // Byte not zero, write it
			*encode++ = *byte, ++code;
        }
		if (!*byte || code == 0xff) { // Input is zero or block completed, restart
			*codep = code, code = 1, codep = encode;
			if (!*byte || length) {
				++encode;
            }
		}
	}
	*codep = code; // Write final code value
	return (size_t)(encode - buffer);
}




/*
//! COBsEncode stuffs "length" bytes of data at the location pointed to by "input"
//! and writes the output to the location pointed to by "output".
//! Returns the number of bytes written to "output".
//! Remove the "restrict" qualifiers if compiling with a pre-C99 C dialect.
//! (copied and adapted from https://github.com/jacquesf/COBS-Consistent-Overhead-Byte-Stuffing/blob/master/cobs.c)
size_t COBSEncode( void * restrict output, const void * restrict input, size_t length){
    unsigned read_index = 0;
    unsigned write_index = 1;
    unsigned code_index = 0;
    uint8_t* out = output;
    const uint8_t* in = input;
    uint8_t code = 1;
    while(read_index < length)
    {
        if(in[read_index] == 0)
        {
            out[code_index] = code;
            code = 1;
            code_index = write_index++;
            read_index++;
        }
        else
        {
            out[write_index++] = in[read_index++];
            code++;
            if(code == 0xFF)
            {
                out[code_index] = code;
                code = 1;
                code_index = write_index++;
            }
        }
    }
    out[code_index] = code;
    return write_index;
}


// copied from https://github.com/datgeezus/cobs
//void cobs_encode(const uint8_t* input, size_t length, uint8_t* output)
size_t CobsEncode(void * out, const void * in, size_t length)
{
    const uint8_t * input = in;
    uint8_t* output = out;
    uint8_t* start = output;
    const uint8_t* end = input + length;    // end of input buffer
    uint8_t* codeBytePos = output;          // holds pointer to the position of a code byte
    uint8_t codeByte = 1;                   // code byte

    // first byte of the input buffer is in the second position of output buffer 
    ++output;

    // iterate over the input bytes
    while (input < end)
    {
        if (*input == 0)
        {
            *codeBytePos = codeByte;        // put the code byte in the corresponding position
            codeByte = 1;                   // reset the code count
            codeBytePos = output;           // update the code byte position
        }
        else
        {
            *output = *input;               // put the
            ++codeByte;                     //
            if (codeByte == 0xFF)
            {
                *codeBytePos = codeByte;
                codeByte = 1;
                codeBytePos = output;
            }
        }
        ++output;
        ++input;
    }

    *codeBytePos = codeByte;
    return output - start;
}


// from https://github.com/TotalKrill/cobs

#define FinishBlock(X) (*code_ptr = (X), code_ptr = dst++, code = 0x01)

//void cobsEncode(const uint8_t *ptr, uint32_t length, uint8_t *dst)
size_t CObsEncode( void * out, const void * in, size_t length )
{
    const uint8_t *ptr = in;
    uint8_t * dst = out;
    uint8_t * start = out;

    const unsigned char *end = ptr + length;
    unsigned char *code_ptr = dst++;
    unsigned char code = 0x01;

    while (ptr < end)
    {
        if (*ptr == 0){
            FinishBlock(code);
        }
        else{
            *dst++ = *ptr;
            if (++code == 0xFF)
                FinishBlock(code);
        }
        ptr++;
    }
    *dst = 0;
    FinishBlock(code);
    return dst-start-1;

}
*/

//! \file cobs.h
//! \author Thomas.Hoehenleitner [at] seerose.net


#ifndef COBS_H_
#define COBS_H_

#include <stddef.h>
#include <stdint.h>

//  size_t COBSEncode(void * restrict output, const void * restrict input, size_t length);
//  size_t CobsEncode(void * out, const void * in, size_t length);
//  size_t CObsEncode(void * out, const void * in, size_t length );
size_t COBSEncode( void * restrict output, const void * restrict data, size_t length);

#endif

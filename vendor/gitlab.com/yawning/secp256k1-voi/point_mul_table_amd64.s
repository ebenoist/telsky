// Copyright (c) 2023 Yawning Angel
//
// SPDX-License-Identifier: BSD-3-Clause

// Code generated by command: go run gen_table_amd64.go. DO NOT EDIT.

//go:build amd64 && !purego

#include "textflag.h"

// func lookupProjectivePoint(tbl *projectivePointMultTable, out *Point, idx uint64)
// Requires: SSE2
TEXT ·lookupProjectivePoint(SB), NOSPLIT|NOFRAME, $0-24
	// This is nice and easy, since it is 3x32-bytes, so this fits
	// neatly into the x86 SIMD registers.  While AVX is significantly
	// nicer to work with, this is done with SSE2 so that only one
	// version of this needs to be implemented.
	// 
	// x0 = x[0] x1 = x[1]
	// y0 = y[0] y1 = y[1]
	// z0 = z[0] z1 = z[1]
	MOVQ   idx+16(FP), AX
	MOVD   AX, X0
	PSHUFD $0x00, X0, X0
	MOVQ   tbl+0(FP), CX

	// Implicit entry tbl[0] = Identity (0, 1, 0)
	PXOR    X1, X1
	PXOR    X2, X2
	PXOR    X3, X3
	PCMPEQL X0, X1
	MOVQ    $0x00000001000003d1, AX
	MOVQ    AX, X4
	PAND    X1, X4
	PXOR    X5, X5
	PXOR    X6, X6
	PXOR    X7, X7

	// For i = 1; i <= 15; i++
	MOVQ $0x0000000000000001, AX

projectiveLookupLoop:
	MOVD    AX, X1
	INCQ    AX
	PSHUFD  $0x00, X1, X1
	PCMPEQL X0, X1
	MOVOU   (CX), X8
	MOVOU   16(CX), X9
	MOVOU   32(CX), X10
	MOVOU   48(CX), X11
	MOVOU   64(CX), X12
	MOVOU   80(CX), X13
	ADDQ    $0x68, CX
	CMPQ    AX, $0x0f
	PAND    X1, X8
	PAND    X1, X9
	PAND    X1, X10
	PAND    X1, X11
	PAND    X1, X12
	PAND    X1, X13
	POR     X8, X2
	POR     X9, X3
	POR     X10, X4
	POR     X11, X5
	POR     X12, X6
	POR     X13, X7
	JLE     projectiveLookupLoop

	// Write out the result.
	MOVQ  out+8(FP), AX
	MOVOU X2, (AX)
	MOVOU X3, 16(AX)
	MOVOU X4, 32(AX)
	MOVOU X5, 48(AX)
	MOVOU X6, 64(AX)
	MOVOU X7, 80(AX)
	RET

// func lookupAffinePoint(tbl *affinePointMultTable, out *affinePoint, idx uint64)
// Requires: SSE2
TEXT ·lookupAffinePoint(SB), NOSPLIT|NOFRAME, $0-24
	// This is nice and easy, since it is 2x32-bytes, so this fits
	// neatly into the x86 SIMD registers.  While AVX is significantly
	// nicer to work with, this is done with SSE2 so that only one
	// version of this needs to be implemented.
	// 
	// x0 = x[0] x1 = x[1]
	// y0 = y[0] y1 = y[1]
	MOVQ   idx+16(FP), AX
	MOVD   AX, X0
	PSHUFD $0x00, X0, X0
	MOVQ   tbl+0(FP), CX

	// Skip idx = 0, addition formula is invalid.
	PXOR X2, X2
	PXOR X3, X3
	PXOR X4, X4
	PXOR X5, X5

	// For i = 1; i <= 15; i++
	MOVQ $0x0000000000000001, AX

affineLookupLoop:
	MOVD    AX, X1
	INCQ    AX
	PSHUFD  $0x00, X1, X1
	PCMPEQL X0, X1
	MOVOU   (CX), X6
	MOVOU   16(CX), X7
	MOVOU   32(CX), X8
	MOVOU   48(CX), X9
	ADDQ    $0x40, CX
	CMPQ    AX, $0x0f
	PAND    X1, X6
	PAND    X1, X7
	PAND    X1, X8
	PAND    X1, X9
	POR     X6, X2
	POR     X7, X3
	POR     X8, X4
	POR     X9, X5
	JLE     affineLookupLoop

	// Write out the result.
	MOVQ  out+8(FP), AX
	MOVOU X2, (AX)
	MOVOU X3, 16(AX)
	MOVOU X4, 32(AX)
	MOVOU X5, 48(AX)
	RET
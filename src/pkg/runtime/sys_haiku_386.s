// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// System calls and other sys.stuff for AMD64, SunOS
// /usr/include/sys/syscall.h for syscall numbers.
//

#include "zasm_GOOS_GOARCH.h"
#include "../../cmd/ld/textflag.h"
TEXT runtime·asmsysvicall6(SB),NOSPLIT,$0
// void runtime·asmstdcall(void *c);
//TEXT runtime·asmstdcall(SB),NOSPLIT,$0
	MOVL	c+0(FP), BX

	// SetLastError(0).
	MOVL	$0, 0x34(FS)

	// Copy args to the stack.
	MOVL	SP, BP
	MOVL	libcall_n(BX), CX	// words
	MOVL	CX, AX
	SALL	$2, AX
	SUBL	AX, SP			// room for args
	MOVL	SP, DI
	MOVL	libcall_args(BX), SI
	CLD
	REP; MOVSL

	// Call stdcall or cdecl function.
	// DI SI BP BX are preserved, SP is not
	CALL	libcall_fn(BX)
	MOVL	BP, SP

	// Return result.
	MOVL	c+0(FP), BX
	MOVL	AX, libcall_r1(BX)
	MOVL	DX, libcall_r2(BX)

	// GetLastError().
	//MOVL	0x34(FS), AX
	//MOVL	AX, libcall_err(BX)

	RET
	
// setldt(int entry, int address, int limit)
TEXT runtime·setldt(SB),NOSPLIT,$0
	MOVL	address+4(FP), CX
	MOVL	CX, 0x14(FS)
	RET

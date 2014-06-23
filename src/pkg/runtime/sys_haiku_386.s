// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// System calls and other sys.stuff for AMD64, SunOS
// /usr/include/sys/syscall.h for syscall numbers.
//

#include "zasm_GOOS_GOARCH.h"
#include "../../cmd/ld/textflag.h"

// void libc·miniterrno(void *(*___errno)(void));
//
// Set the TLS errno pointer in M.
//
// Called using runtime·asmcgocall from os_solaris.c:/minit.
TEXT runtime·miniterrno(SB),NOSPLIT,$0
	// FIXME: implement
	RET
	MOVL	4(SP), DI
	CALL	DI	// SysV ABI so returns in AX
	get_tls(CX)
	MOVL	m(CX), BX
	MOVL	AX,	m_perrno(BX)
	RET


TEXT runtime·asmsysvicall6(SB),NOSPLIT,$0
// void runtime·asmstdcall(void *c);
//TEXT runtime·asmstdcall(SB),NOSPLIT,$0
	MOVL	c+0(FP), BX

	// SetLastError(0).
	//MOVL	$0, 0x34(FS)

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
	MOVL	$0, AX
	MOVL	AX, libcall_err(BX)
	//FIXME: implement this

	RET

// uint32 tstart_sysvicall(M *newm);
TEXT runtime·tstart_sysvicall(SB),NOSPLIT,$0
	MOVL	newm+4(SP), CX		// m
	MOVL	m_g0(CX), DX		// g

	// Make TLS entries point at g and m.
	//get_tls(BX)
	LEAL	m_tls(CX), BX
	MOVL	BX, 0x14(FS) //TODO: Fix this so that I'm not hardcoding a TLS slot
	MOVL	CX, m(BX)
	MOVL	DX, g(BX)

	// Layout new m scheduler stack on os stack.
	MOVL	SP, AX
	MOVL	AX, g_stackbase(DX)
	SUBL	$(64*1024), AX		// stack size
	MOVL	AX, g_stackguard(DX)
	MOVL	AX, g_stackguard0(DX)

	// Someday the convention will be D is always cleared.
	CLD

	CALL	runtime·stackcheck(SB)	// clobbers AX,CX
	CALL	runtime·mstart(SB)

	XORL	AX, AX			// return 0 == success

	RET
	
// setldt(int entry, int address, int limit)
TEXT runtime·setldt(SB),NOSPLIT,$0
	MOVL	address+4(FP), CX
	MOVL	CX, 0x14(FS)
	RET

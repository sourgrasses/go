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
	MOVL	4(SP), DI
	CALL	DI	// SysV ABI so returns in AX
	get_tls(CX)
	MOVL	m(CX), BX
	MOVL	AX,	m_perrno(BX)
	RET


TEXT runtime·asmsysvicall6(SB),NOSPLIT,$0
// void runtime·asmstdcall(void *c);
//TEXT runtime·asmstdcall(SB),NOSPLIT,$0
	get_tls(CX)
	MOVL	m(CX), BX
	MOVL	m_perrno(BX), DX
	CMPL	DX, $0
	JEQ	skiperrno1
	// reset errno to 0
	MOVL	$0, 0(DX)
skiperrno1:
	MOVL	c+0(FP), BX
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

	// copy over errno
	get_tls(CX)
	MOVL	m(CX), AX
	MOVL	m_perrno(AX), AX
	CMPL	AX, $0
	JEQ	skiperrno2
	MOVL	0(AX), AX
	MOVL	AX, libcall_err(BX)
skiperrno2:
	RET

// uint32 tstart_sysvicall(M *newm);
TEXT runtime·tstart_sysvicall(SB),NOSPLIT,$0
	MOVL	newm+4(SP), CX		// m
	MOVL	m_g0(CX), DX		// g

	// Make TLS entries point at g and m.
	//get_tls(BX)
	LEAL	m_tls(CX), BX
	MOVL	BX, 0xfc(FS) //TODO: Fix this so that I'm not hardcoding a TLS slot
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
	MOVL	CX, 0xfc(FS)
	RET

TEXT runtime·sigtramp(SB),NOSPLIT,$88
	// 24 + 20 + 24 + 4 + 8 = 80 bytes of stack space
	// can use: ax, cx, dx
	// libcall = 5 * 4 = 20
	// scratch = 6 * 4 = 24 (44, 48, 52, 56, 60, 64)
	// errno = 1*4 (68)
	// Registers trampled on: bx, di = 8 (72, 76)
	// Regsters that are not used, but needs save:
	// bp, si = 8 ( 80, 84)
	MOVL	BX, 72(SP)
	MOVL	DI, 76(SP)
	MOVL	BP, 80(SP)
	MOVL	SI, 84(SP)

	get_tls(CX)

	// check that m exists
	MOVL	m(CX), BX
	CMPL	BX, $0
	JNE	6(PC)
	MOVL	signo+0(FP), BX
	MOVL	BX, 0(SP)
	MOVL	$runtime·badsignal(SB), AX
	CALL	AX
	JMP 	sigtramp_ret

	// save g
	MOVL	g(CX), DI
	MOVL	DI, 20(SP)
	// save libcall (as Solaris's port does)
	LEAL	m_libcall(BX), DI
	MOVL	libcall_fn(DI), AX
	MOVL	AX, 24(SP)
	MOVL	libcall_args(DI), AX
	MOVL	AX, 28(SP)
	MOVL	libcall_n(DI), AX
	MOVL	AX, 32(SP)
	MOVL	libcall_r1(DI), AX
	MOVL	AX, 36(SP)
	MOVL	libcall_r2(DI), AX
	MOVL	AX, 40(SP)
	// save scratch
	LEAL	m_scratch(BX), DI
	MOVL	0(DI), AX
	MOVL	AX, 44(SP)
	MOVL	4(DI), AX
	MOVL	AX, 48(SP)
	MOVL	8(DI), AX
	MOVL	AX, 52(SP)
	MOVL	12(DI), AX
	MOVL	AX, 56(SP)
	MOVL	16(DI), AX
	MOVL	AX, 60(SP)
	MOVL	20(DI), AX
	MOVL	AX, 64(SP)
	// save errno
//	MOVL	m_perrno(BX), DI
//	MOVL	0(DI), DI
//	MOVL	DI, 68(SP)

	// g = m->gsignal
	MOVL	m_gsignal(BX), BX
	MOVL	BX, g(CX)

	// copy arguments for call to sighandler
	MOVL	signo+0(FP), BX
	MOVL	BX, 0(SP)
	MOVL	info+4(FP), BX
	MOVL	BX, 4(SP)
	MOVL	context+8(FP), BX
	MOVL	BX, 8(SP)
	MOVL	DI, 12(SP)

	CALL	runtime·sighandler(SB)

	// reload TLS and M
	get_tls(CX)
	MOVL	m(CX), BX

	// restore g
	get_tls(CX)
	MOVL	20(SP), BX
	MOVL	BX, g(CX)
	// restore libcall
	LEAL	m_libcall(BX), DI
	MOVL	24(SP), AX
	MOVL	AX, libcall_fn(DI)
	MOVL	28(SP), AX
	MOVL	AX, libcall_args(DI)
	MOVL	32(SP), AX
	MOVL	AX, libcall_n(DI)
	MOVL	36(SP), AX
	MOVL	AX, libcall_r1(DI)
	MOVL	40(SP), AX
	MOVL	AX, libcall_r2(DI)

	// restore scratch
	LEAL	m_scratch(BX), DI
	MOVL	44(SP), AX
	MOVL	AX, 0(DI)
	MOVL	48(SP), AX
	MOVL	AX, 4(DI)
	MOVL	52(SP), AX
	MOVL	AX, 8(DI)
	MOVL	56(SP), AX
	MOVL	AX, 12(DI)
	MOVL	60(SP), AX
	MOVL	AX, 16(DI)
	MOVL	64(SP), AX
	MOVL	AX, 20(DI)

	// restore errno
//	MOVL	m_perrno(BX), DI
//	MOVL	68(SP), AX
//	MOVL	AX, 0(DI)
sigtramp_ret:
	// restore registers
	MOVL	72(SP), BX
	MOVL	76(SP), DI
	MOVL	80(SP), BP
	MOVL	84(SP), SI
	RET

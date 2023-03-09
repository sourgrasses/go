// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT ·Syscall(SB),NOSPLIT,$0
	JMP	runtime·syscall_syscall(SB)

TEXT ·RawSyscall(SB),NOSPLIT,$0
	JMP	runtime·syscall_rawsyscall(SB)

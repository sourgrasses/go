// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

type sigTabT struct {
	flags int32
	name  string
}

var sigtable = [...]sigTabT{
	/* 0 */ {0, "SIGNONE: no trap"},
	/* 1 */ {_SigNotify + _SigKill, "SIGHUP: hangup"},
	/* 2 */ {_SigNotify + _SigKill, "SIGINT: interrupt"},
	/* 3 */ {_SigNotify + _SigThrow, "SIGQUIT: quit"},
	/* 4 */ {_SigThrow + _SigUnblock, "SIGILL: illegal instruction"},
	/* 5 */ {_SigThrow + _SigUnblock, "SIGCHLD: child process exited"},
	/* 6 */ {_SigNotify + _SigThrow, "SIGABRT: abort"},
	/* 7 */ {_SigNotify, "SIGPIPE: broken pipe"},
	/* 8 */ {_SigPanic + _SigUnblock, "SIGFPE: floating point exception"},
	/* 9 */ {0, "SIGKILL: killed (by death)"},
	/* 10 */ {_SigNotify + _SigDefault, "SIGSTOP: stopped"},
	/* 11 */ {_SigPanic + _SigUnblock, "SIGSEGV: segmentation violation"},
	/* 12 */ {0, "SIGCONT: continued"},
	/* 13 */ {0, "SIGTSTP: stopped (tty output)"},
	/* 14 */ {_SigNotify, "SIGALRM: alarm"},
	/* 15 */ {_SigNotify + _SigKill, "SIGTERM: termination requested"},
	/* 16 */ {_SigNotify + _SigDefault, "SIGTTIN: stopped (tty input)"},
	/* 17 */ {_SigNotify + _SigDefault, "SIGTTOU: stopped (tty output)"},
	/* 18 */ {_SigNotify, "SIGUSR1: user defined signal 1"},
	/* 19 */ {_SigNotify, "SIGUSR2: user defined signal 2"},
	/* 20 */ {_SigNotify, "SIGWINCH: window size change"},
	/* 21 */ {_SigNotify + _SigKill, "SIGKILLTHR: kill thread"},
	/* 22 */ {_SigThrow, "SIGTRAP: trace/breakpoint trap"},
	/* 23 */ {_SigNotify, "SIGPOLL: pollable event"},
	/* 24 */ {_SigNotify, "SIGPROF: profiling timer expired"},
	/* 25 */ {_SigNotify, "SIGSYS: bad system call"},
	/* 26 */ {_SigNotify, "SIGURG: high bandwidth data is available at socket"},
	/* 27 */ {_SigNotify, "SIGVTALRM: virtual timer expired"},
	/* 28 */ {_SigNotify, "SIGXCPU: CPU time limit expired"},
	/* 29 */ {_SigNotify, "SIGXFSZ: file size limit expired"},
	/* 30 */ {_SigPanic, "SIGBUS: bus error"},
	/* 31 */ {_SigNotify, "SIGRESERVED1: reserved signal 1"},
	/* 32 */ {_SigNotify, "SIGRESERVED2: reserved signal 2"},
}

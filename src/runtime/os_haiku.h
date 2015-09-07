// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#define SS_DISABLE 2

// Linux-specific system calls
//int32	runtime·futex(uint32*, int32, uint32, Timespec*, uint32*, uint32);
//int32	runtime·clone(int32, void*, M*, G*, void(*)(void));

struct Sigaction;
void	runtime·sigaction(int32, struct Sigaction*, struct Sigaction*);

void	runtime·sigaltstack(Sigaltstack*, Sigaltstack*);
void	runtime·sigpanic(void);
void runtime·setitimer(int32, Itimerval*, Itimerval*);


#define	NSIG	65
#define	SI_USER 0

void	runtime·sigprocmask(int32, Sigset*, Sigset*);
void	runtime·unblocksignals(void);
#define SIG_SETMASK 3

#define RLIMIT_AS 6
typedef struct Rlimit Rlimit;
struct Rlimit {
	uintptr	rlim_cur;
	uintptr	rlim_max;
};
int32	runtime·getrlimit(int32, Rlimit*);

// Call a library function with SysV conventions,
// and switch to os stack during the call.
#pragma	varargck	countpos	runtime·sysvicall6	2
#pragma	varargck	type		runtime·sysvicall6	uintptr
#pragma	varargck	type		runtime·sysvicall6	int32
void	runtime·asmsysvicall6(void *c);
uintptr	runtime·sysvicall6(uintptr fn, int32 count, ...);

void	runtime·miniterrno(void *fn);

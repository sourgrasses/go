// Created by cgo -cdefs - DO NOT EDIT
// cgo -cdefs defs_haiku.go


enum {
	EINTR	= -0x7ffffff6,
	EFAULT	= -0x7fffecff,
	EAGAIN	= -0x7ffffff5,
	ENOMEM	= -0x80000000,

	PROT_NONE	= 0x0,
	PROT_READ	= 0x1,
	PROT_WRITE	= 0x2,
	PROT_EXEC	= 0x4,

	MAP_ANON	= 0x8,
	MAP_PRIVATE	= 0x2,
	MAP_FIXED	= 0x4,

	SA_SIGINFO	= 0x40,
	SA_RESTART	= 0x10,
	SA_ONSTACK	= 0x20,

	SIGHUP		= 0x1,
	SIGINT		= 0x2,
	SIGQUIT		= 0x3,
	SIGILL		= 0x4,
	SIGTRAP		= 0x16,
	SIGABRT		= 0x6,
	SIGFPE		= 0x8,
	SIGKILL		= 0x9,
	SIGBUS		= 0x1e,
	SIGSEGV		= 0xb,
	SIGSYS		= 0x19,
	SIGPIPE		= 0x7,
	SIGALRM		= 0xe,
	SIGTERM		= 0xf,
	SIGURG		= 0x1a,
	SIGSTOP		= 0xa,
	SIGTSTP		= 0xd,
	SIGCONT		= 0xc,
	SIGCHLD		= 0x5,
	SIGTTIN		= 0x10,
	SIGTTOU		= 0x11,
	SIGXCPU		= 0x1c,
	SIGXFSZ		= 0x1d,
	SIGVTALRM	= 0x1b,
	SIGPROF		= 0x18,
	SIGWINCH	= 0x14,
	SIGUSR1		= 0x12,
	SIGUSR2		= 0x13,

	FPE_INTDIV	= 0x14,
	FPE_INTOVF	= 0x15,
	FPE_FLTDIV	= 0x16,
	FPE_FLTOVF	= 0x17,
	FPE_FLTUND	= 0x18,
	FPE_FLTRES	= 0x19,
	FPE_FLTINV	= 0x1a,
	FPE_FLTSUB	= 0x1b,

	BUS_ADRALN	= 0x28,
	BUS_ADRERR	= 0x29,
	BUS_OBJERR	= 0x2a,

	SEGV_MAPERR	= 0x1e,
	SEGV_ACCERR	= 0x1f,

	ITIMER_REAL	= 0x1,
	ITIMER_VIRTUAL	= 0x2,
	ITIMER_PROF	= 0x3,

	PTHREAD_CREATE_DETACHED	= 0x1,

	HOST_NAME_MAX	= 0xff,

	O_NONBLOCK	= 0x80,
	FD_CLOEXEC	= 0x1,
	F_GETFL		= 0x8,
	F_SETFL		= 0x10,
	F_SETFD		= 0x4,
};

typedef struct SemT SemT;
typedef struct StackT StackT;
typedef struct Siginfo Siginfo;
typedef struct Sigaction Sigaction;
typedef struct ExtendedRegs ExtendedRegs;
typedef struct Mcontext Mcontext;
typedef struct Ucontext Ucontext;
typedef struct Timespec Timespec;
typedef struct Timeval Timeval;
typedef struct Itimerval Itimerval;
typedef struct Stat Stat;

#pragma pack on

struct SemT {
	int32	id;
	int32	_padding[3];
};

typedef	uint64	Sigset;
struct StackT {
	byte	*ss_sp;
	uint32	ss_size;
	int32	ss_flags;
};
typedef	StackT	Sigaltstack;

struct Siginfo {
	int32	si_signo;
	int32	si_code;
	int32	si_errno;
	int32	si_pid;
	uint32	si_uid;
	byte	*si_addr;
	int32	si_status;
	int32	si_band;
	byte	si_value[4];
};
struct Sigaction {
	byte	anon0[4];
	uint64	sa_mask;
	int32	sa_flags;
	byte	*sa_userdata;
};

struct ExtendedRegs {
	byte	state[512];
	uint32	format;
};

struct Mcontext {
	uint32	eip;
	uint32	eflags;
	uint32	eax;
	uint32	ecx;
	uint32	edx;
	uint32	esp;
	uint32	ebp;
	uint32	_reserved_1;
	ExtendedRegs	xregs;
	uint32	edi;
	uint32	esi;
	uint32	ebx;
};
struct Ucontext {
	Ucontext	*uc_link;
	uint64	uc_sigmask;
	StackT	uc_stack;
	Mcontext	uc_mcontext;
};

struct Timespec {
	int32	tv_sec;
	int32	tv_nsec;
};
struct Timeval {
	int32	tv_sec;
	int32	tv_usec;
};
struct Itimerval {
	Timeval	it_interval;
	Timeval	it_value;
};

typedef	void	*Pthread;
typedef	void	*PthreadAttr;

struct Stat {
	int32	st_dev;
	int64	st_ino;
	uint32	st_mode;
	int32	st_nlink;
	uint32	st_uid;
	uint32	st_gid;
	int64	st_size;
	int32	st_rdev;
	int32	st_blksize;
	Timespec	st_atim;
	Timespec	st_mtim;
	Timespec	st_ctim;
	Timespec	st_crtim;
	uint32	st_type;
	int64	st_blocks;
};


#pragma pack off

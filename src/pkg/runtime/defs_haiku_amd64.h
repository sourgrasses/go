// Created by cgo -cdefs - DO NOT EDIT
// cgo -cdefs defs_solaris.go defs_solaris_amd64.go


enum {
	EINTR		= -0x7ffffff6,
	EFAULT		= -0x7fffecff,
	EAGAIN		= -0x7ffffff5,
	ENOMEM		= -0x80000000,
	ETIMEDOUT	= -0x7ffffff7,
	EACCES		= -0x7ffffffe,

	PROT_NONE	= 0x0,
	PROT_READ	= 0x1,
	PROT_WRITE	= 0x2,
	PROT_EXEC	= 0x4,

	MAP_ANON	= 0x8,
	MAP_PRIVATE	= 0x2,
	MAP_FIXED	= 0x4,

	MADV_FREE	= 0x5,

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

	FORK_NOSIGCHLD	= 0x1,
	FORK_WAITPID	= 0x2,

	HOST_NAME_MAX	= 0xff,

	O_NONBLOCK	= 0x80,
	FD_CLOEXEC	= 0x1,
	F_GETFL		= 0x8,
	F_SETFL		= 0x10,
	F_SETFD		= 0x4,

	_SC_NPROCESSORS_ONLN	= 0x23,

	B_PAGE_SIZE	= 0x1000,

	B_ERROR	= -0x1,
};

typedef struct SemT SemT;
typedef struct Sigaltstack Sigaltstack;
typedef struct StackT StackT;
typedef struct Siginfo Siginfo;
typedef struct Sigaction Sigaction;
typedef struct Fpregset Fpregset;
typedef struct Mcontext Mcontext;
typedef struct Ucontext Ucontext;
typedef struct Timespec Timespec;
typedef struct Timeval Timeval;
typedef struct Itimerval Itimerval;
typedef struct PortEvent PortEvent;
typedef struct PthreadAttr PthreadAttr;
typedef struct Stat Stat;

#pragma pack on

struct SemT {
	uint32	sem_count;
	uint16	sem_type;
	uint16	sem_magic;
	uint64	sem_pad1[3];
	uint64	sem_pad2[2];
};

struct Sigaltstack {
	byte	*ss_sp;
	uint64	ss_size;
	int32	ss_flags;
	byte	Pad_cgo_0[4];
};
typedef	uint64	Sigset;
struct StackT {
	byte	*ss_sp;
	uint64	ss_size;
	int32	ss_flags;
	byte	Pad_cgo_0[4];
};

struct Siginfo {
	int32	si_signo;
	int32	si_code;
	int32	si_errno;
	int32	si_pad;
	byte	__data[240];
};
struct Sigaction {
	byte	anon0[8];
	uint64	sa_mask;
	int32	sa_flags;
	byte	*sa_userdata;
};

struct Fpregset {
	byte	fp_reg_set[528];
};
struct Mcontext {
	int64	gregs[28];
	Fpregset	fpregs;
};
struct Ucontext {
	uint64	uc_flags;
	Ucontext	*uc_link;
	Sigset	uc_sigmask;
	StackT	uc_stack;
	byte	Pad_cgo_0[8];
	Mcontext	uc_mcontext;
	int64	uc_filler[5];
	byte	Pad_cgo_1[8];
};

struct Timespec {
	int64	tv_sec;
	int64	tv_nsec;
};
struct Timeval {
	int64	tv_sec;
	int64	tv_usec;
};
struct Itimerval {
	Timeval	it_interval;
	Timeval	it_value;
};

struct PortEvent {
	int32	portev_events;
	uint16	portev_source;
	uint16	portev_pad;
	uint64	portev_object;
	byte	*portev_user;
};
typedef	uint32	Pthread;
struct PthreadAttr {
	byte	*__pthread_attrp;
};

struct Stat {
	uint64	st_dev;
	uint64	st_ino;
	uint32	st_mode;
	uint32	st_nlink;
	uint32	st_uid;
	uint32	st_gid;
	uint64	st_rdev;
	int64	st_size;
	Timespec	st_atim;
	Timespec	st_mtim;
	Timespec	st_ctim;
	int32	st_blksize;
	byte	Pad_cgo_0[4];
	int64	st_blocks;
	int8	st_fstype[16];
};


#pragma pack off
// Created by cgo -cdefs - DO NOT EDIT
// cgo -cdefs defs_solaris.go defs_solaris_amd64.go


enum {
	REG_RDI		= 0x8,
	REG_RSI		= 0x9,
	REG_RDX		= 0xc,
	REG_RCX		= 0xd,
	REG_R8		= 0x7,
	REG_R9		= 0x6,
	REG_R10		= 0x5,
	REG_R11		= 0x4,
	REG_R12		= 0x3,
	REG_R13		= 0x2,
	REG_R14		= 0x1,
	REG_R15		= 0x0,
	REG_RBP		= 0xa,
	REG_RBX		= 0xb,
	REG_RAX		= 0xe,
	REG_GS		= 0x17,
	REG_FS		= 0x16,
	REG_ES		= 0x18,
	REG_DS		= 0x19,
	REG_TRAPNO	= 0xf,
	REG_ERR		= 0x10,
	REG_RIP		= 0x11,
	REG_CS		= 0x12,
	REG_RFLAGS	= 0x13,
	REG_RSP		= 0x14,
	REG_SS		= 0x15,
};


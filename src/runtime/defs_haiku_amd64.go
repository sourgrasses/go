// Created by cgo -cdefs - DO NOT EDIT
// cgo -cdefs defs_haiku.go

package runtime

const (
	EINTR		= 0x8000000a
	EFAULT		= 0x80001301
	EAGAIN		= 0x8000000b
	ENOMEM		= 0x80000000
	ETIMEDOUT	= 0x80000009
	EACCES		= 0x80000002

	_EINTR		= -0x7ffffff6
	_EFAULT		= -0x7fffecff
	_EAGAIN		= -0x7ffffff5
	_ENOMEM		= -0x80000000
	_ETIMEDOUT	= -0x7ffffff7
	_EACCES		= -0x7ffffffe

	PROT_NONE	= 0x0
	PROT_READ	= 0x1
	PROT_WRITE	= 0x2
	PROT_EXEC	= 0x4

	MAP_ANON	= 0x8
	MAP_PRIVATE	= 0x2
	MAP_FIXED	= 0x4

	_SA_SIGINFO	= 0x40
	_SA_RESTART	= 0x10
	_SA_ONSTACK	= 0x20

	_SIGHUP		= 0x1
	_SIGINT		= 0x2
	_SIGQUIT		= 0x3
	_SIGILL		= 0x4
	_SIGTRAP		= 0x16
	_SIGABRT		= 0x6
	_SIGFPE		= 0x8
	_SIGKILL		= 0x9
	_SIGBUS		= 0x1e
	_SIGSEGV		= 0xb
	_SIGSYS		= 0x19
	_SIGPIPE		= 0x7
	_SIGALRM		= 0xe
	_SIGTERM		= 0xf
	_SIGURG		= 0x1a
	_SIGSTOP		= 0xa
	_SIGTSTP		= 0xd
	_SIGCONT		= 0xc
	_SIGCHLD		= 0x5
	_SIGTTIN		= 0x10
	_SIGTTOU		= 0x11
	_SIGXCPU		= 0x1c
	_SIGXFSZ		= 0x1d
	_SIGVTALRM	= 0x1b
	_SIGPROF		= 0x18
	_SIGWINCH	= 0x14
	_SIGUSR1		= 0x12
	_SIGUSR2		= 0x13

	_FPE_INTDIV	= 0x14
	_FPE_INTOVF	= 0x15
	_FPE_FLTDIV	= 0x16
	_FPE_FLTOVF	= 0x17
	_FPE_FLTUND	= 0x18
	_FPE_FLTRES	= 0x19
	_FPE_FLTINV	= 0x1a
	_FPE_FLTSUB	= 0x1b

	_BUS_ADRALN	= 0x28
	_BUS_ADRERR	= 0x29
	_BUS_OBJERR	= 0x2a

	_SEGV_MAPERR	= 0x1e
	_SEGV_ACCERR	= 0x1f

	_ITIMER_REAL	= 0x1
	_ITIMER_VIRTUAL	= 0x2
	_ITIMER_PROF	= 0x3

	_PTHREAD_CREATE_DETACHED	= 0x1

	HOST_NAME_MAX	= 0xff

	O_NONBLOCK	= 0x80
	FD_CLOEXEC	= 0x1
	F_GETFL		= 0x8
	F_SETFL		= 0x10
	F_SETFD		= 0x4

	_SC_NPROCESSORS_ONLN	= 0x23

	B_PAGE_SIZE	= 0x1000

	B_ERROR	= -0x1
)

type pthread struct {
	__pthread_p *byte
}
type pthreadattr struct {
	__pthread_attrp *byte
}

type semt struct {
	sem_id   int32
	padding  [3]int32
}

type sigset uint64;

type sigaltstackt struct {
	ss_sp       uintptr
	ss_size     uint64
	ss_flags    int32
	padding     uint32
}

type stackt struct {
	ss_sp       uintptr
	ss_size     uint64
	ss_flags    int32
	padding     uint32
}

type siginfo struct {
	si_signo    int32
	si_code     int32
	si_errno    int32
	si_pid      int32
	si_uid      uint32
	padding0    uint32
	si_addr     uint64
	si_status   int32
	padding1	uint32
	si_band     int64
	si_value    [8]byte
}

type sigactiont struct {
	sa_handler  uintptr
	sa_mask     sigset
	sa_flags    int32
	padding    uint32
	sa_userdata uintptr
}

type xmmregs struct {
	xmm0 [16]uint8
	xmm1 [16]uint8
	xmm2 [16]uint8
	xmm3 [16]uint8
	xmm4 [16]uint8
	xmm5 [16]uint8
	xmm6 [16]uint8
	xmm7 [16]uint8
	xmm8 [16]uint8
	xmm9 [16]uint8
	xmm10 [16]uint8
	xmm11 [16]uint8
	xmm12 [16]uint8
	xmm13 [16]uint8
	xmm14 [16]uint8
	xmm15 [16]uint8
}

type fpustate struct {
	control    uint16
	status     uint16
	tag        uint16
	opcode     uint16
	rip        uint64
	rdp        uint64
	mxcsr      uint32
	mscsr_mask uint32
	anon0      [128]byte
	xmm        xmmregs
	reserved_416_511 [96]uint8
}

type mcontext struct {
	rax     uint64
	rbx     uint64
	rcx     uint64
	rdx     uint64
	rdi     uint64
	rsi     uint64
	rbp     uint64
	r8      uint64
	r9      uint64
	r10     uint64
	r11     uint64
	r12     uint64
	r13     uint64
	r14     uint64
	r15     uint64
	rsp     uint64
	rip     uint64
	rflags  uint64
	fpu 	fpustate
}

type ucontext struct {
	uc_link     *ucontext
	uc_sigmask  uint64
	uc_stack    stackt
	uc_mcontext mcontext
}

type timespec struct {
	tv_sec  int32
	padding uint32
	tv_nsec int64
}

type timeval struct {
	tv_sec  int32
	tv_usec int32
}

func (tv *timeval) set_usec(x int32) {
	tv.tv_usec = x
}

type itimerval struct {
	it_interval timeval
	it_value    timeval
}


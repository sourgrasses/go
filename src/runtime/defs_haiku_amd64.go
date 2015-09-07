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

	PROT_NONE	= 0x0
	PROT_READ	= 0x1
	PROT_WRITE	= 0x2
	PROT_EXEC	= 0x4

	MAP_ANON	= 0x8
	MAP_PRIVATE	= 0x2
	MAP_FIXED	= 0x4

	HOST_NAME_MAX	= 0xff

	O_NONBLOCK	= 0x80
	FD_CLOEXEC	= 0x1
	F_GETFL		= 0x8
	F_SETFL		= 0x10
	F_SETFD		= 0x4

	B_PAGE_SIZE	= 0x1000

	B_ERROR	= -0x1
)

type pthread struct {
	__pthread_p *byte
}
type pthreadattr struct {
	__pthread_attrp *byte
}

type sigset uint64;

type sigaltstackt struct {
	ss_sp       uintptr
	ss_size     uint64
	ss_flags    int32
	padding     uint32
}

type sigactiont struct {
	sa_handler  uintptr
	sa_mask     sigset
	sa_flags    int32
	padding    uint32
	sa_userdata uintptr
}

func (tv *timeval) set_usec(x int32) {
	tv.tv_usec = x
}


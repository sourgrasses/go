// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

type libcFunc uintptr

var (
	libc__errnop,
	libc_clock_gettime,
	libc_close,
	libc_exit,
	libc_fstat,
	libc_getcontext,
	libc_getrlimit,
	libc_kill,
	libc_malloc,
	libc_mmap,
	libc_munmap,
	libc_open,
	libc_pthread_attr_destroy,
	libc_pthread_attr_getstacksize,
	libc_pthread_attr_init,
	libc_pthread_attr_setdetachstate,
	libc_pthread_attr_setstack,
	libc_pthread_create,
	libc_raise,
	libc_read,
	libc_sched_yield,
	libc_select,
	libc_sem_init,
	libc_sem_post,
	libc_sem_timedwait,
	libc_sem_wait,
	libc_setitimer,
	libc_sigaction,
	libc_sigaltstack,
	libc_sigprocmask,
	libc_sysconf,
	libc_usleep,
	libc_write,
	libc_area_for libcFunc
)

var asmsysvicall6 libcFunc

//go:nosplit
func closefd(fd int32) int32 {
	return int32(sysvicall1(&libc_close, uintptr(fd)))
}



func raiseproc(sig int32) /* int32 */ {
	pid := sysvicall0(&libc_getpid)
	sysvicall2(&libc_kill, pid, uintptr(sig))
}


//go:noescape
func sigfwd(fn uintptr, sig uint32, info *siginfo, ctx unsafe.Pointer)

//go:nosplit
func sysvicall0(fn *libcFunc) uintptr {
	libcall := &getg().m.libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 0
	libcall.args = uintptr(unsafe.Pointer(fn)) // it's unused but must be non-nil, otherwise crashes
	asmcgocall(unsafe.Pointer(&asmsysvicall6), unsafe.Pointer(libcall))
	return libcall.r1
}

//go:nosplit
func sysvicall1(fn *libcFunc, a1 uintptr) uintptr {
	libcall := &getg().m.libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 1
	// TODO(rsc): Why is noescape necessary here and below?
	libcall.args = uintptr(noescape(unsafe.Pointer(&a1)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6), unsafe.Pointer(libcall))
	return libcall.r1
}

//go:nosplit
func sysvicall2(fn *libcFunc, a1, a2 uintptr) uintptr {
	libcall := &getg().m.libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 2
	libcall.args = uintptr(noescape(unsafe.Pointer(&a1)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6), unsafe.Pointer(libcall))
	return libcall.r1
}

//go:nosplit
func sysvicall3(fn *libcFunc, a1, a2, a3 uintptr) uintptr {
	libcall := &getg().m.libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 3
	libcall.args = uintptr(noescape(unsafe.Pointer(&a1)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6), unsafe.Pointer(libcall))
	return libcall.r1
}

//go:nosplit
func sysvicall4(fn *libcFunc, a1, a2, a3, a4 uintptr) uintptr {
	libcall := &getg().m.libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 4
	libcall.args = uintptr(noescape(unsafe.Pointer(&a1)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6), unsafe.Pointer(libcall))
	return libcall.r1
}

//go:nosplit
func sysvicall5(fn *libcFunc, a1, a2, a3, a4, a5 uintptr) uintptr {
	libcall := &getg().m.libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 5
	libcall.args = uintptr(noescape(unsafe.Pointer(&a1)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6), unsafe.Pointer(libcall))
	return libcall.r1
}

//go:nosplit
func sysvicall6(fn *libcFunc, a1, a2, a3, a4, a5, a6 uintptr) uintptr {
	libcall := &getg().m.libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 6
	libcall.args = uintptr(noescape(unsafe.Pointer(&a1)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6), unsafe.Pointer(libcall))
	return libcall.r1
}

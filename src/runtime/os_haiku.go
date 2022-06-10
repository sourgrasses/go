// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

type mts struct {
	tv_sec  int64
	tv_nsec int64
}

type mscratch struct {
	v [6]uintptr
}

type mOS struct {
	waitsema uintptr // semaphore for parking on locks
	perrno   *int32  // pointer to tls errno
	// // these are here because they are too large to be on the stack
	// // of low-level NOSPLIT functions.
	// //LibCall       libcall;
	ts      mts
	scratch mscratch
}

const (
	_SS_DISABLE  = 4
	_NSIG        = 33
	_SI_USER     = 0
	_SIG_SETMASK = 3 // Modified for Haiku support
	_RLIMIT_AS   = 6 // Modified for Haiku support
	_SIG_UNBLOCK = 2

	_CLOCK_REALTIME  = 0xffffffff
)


type libcFunc uintptr

//go:linkname asmsysvicall6x runtime.asmsysvicall6
var asmsysvicall6x libcFunc // name to take addr of asmsysvicall6

func asmsysvicall6() // declared for vet; do NOT call

//go:nosplit
//go:cgo_unsafe_args
func sysvicall0(fn *libcFunc) uintptr {
	var libcall libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 0
	libcall.args = uintptr(unsafe.Pointer(fn)) // it's unused but must be non-nil, otherwise crashes
	asmcgocall(unsafe.Pointer(&asmsysvicall6x), unsafe.Pointer(&libcall))
	return libcall.r1
}

//go:nosplit
//go:cgo_unsafe_args
func sysvicall1(fn *libcFunc, a1 uintptr) uintptr {
	var libcall libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 1
	// TODO(rsc): Why is noescape necessary here and below?
	libcall.args = uintptr(noescape(unsafe.Pointer(&a1)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6x), unsafe.Pointer(&libcall))
	return libcall.r1
}

//go:nosplit
//go:cgo_unsafe_args
func sysvicall2(fn *libcFunc, a1, a2 uintptr) uintptr {
	var libcall libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 2
	libcall.args = uintptr(noescape(unsafe.Pointer(&a1)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6x), unsafe.Pointer(&libcall))
	return libcall.r1
}

//go:nosplit
//go:cgo_unsafe_args
func sysvicall3(fn *libcFunc, a1, a2, a3 uintptr) uintptr {
	var libcall libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 3
	libcall.args = uintptr(noescape(unsafe.Pointer(&a1)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6x), unsafe.Pointer(&libcall))
	return libcall.r1
}

//go:nosplit
//go:cgo_unsafe_args
func sysvicall4(fn *libcFunc, a1, a2, a3, a4 uintptr) uintptr {
	var libcall libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 4
	libcall.args = uintptr(noescape(unsafe.Pointer(&a1)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6x), unsafe.Pointer(&libcall))
	return libcall.r1
}

//go:nosplit
//go:cgo_unsafe_args
func sysvicall5(fn *libcFunc, a1, a2, a3, a4, a5 uintptr) uintptr {
	var libcall libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 5
	libcall.args = uintptr(noescape(unsafe.Pointer(&a1)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6x), unsafe.Pointer(&libcall))
	return libcall.r1
}

//go:nosplit
//go:cgo_unsafe_args
func sysvicall6(fn *libcFunc, a1, a2, a3, a4, a5, a6 uintptr) uintptr {
	var libcall libcall
	libcall.fn = uintptr(unsafe.Pointer(fn))
	libcall.n = 6
	libcall.args = uintptr(noescape(unsafe.Pointer(&a1)))
	asmcgocall(unsafe.Pointer(&asmsysvicall6x), unsafe.Pointer(&libcall))
	return libcall.r1
}

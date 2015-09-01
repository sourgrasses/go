// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

// NOTE: vec must be just 1 byte long here.
// Mincore returns ENOMEM if any of the pages are unmapped,
// but we want to know that all of the pages are unmapped.
// To make these the same, we can only ask about one page
// at a time. See golang.org/issue/7476.
var addrspace_vec [1]byte

func addrspace_free(v unsafe.Pointer, n uintptr) bool {
	var chunk uintptr
	for off := uintptr(0); off < n; off += chunk {
		chunk = B_PAGE_SIZE
		if chunk > (n - off) {
			chunk = n - off
		}
		errval := int32(sysvicall1(&libc_area_for, uintptr(v)+off))
		// B_ERROR means unmapped, which is what we want.
		// Anything else we assume means the pages are mapped.
		if errval != B_ERROR {
			return false
		}
	}
	return true
}

func mmap_fixed(v unsafe.Pointer, n uintptr, prot, flags, fd int32, offset uint32) unsafe.Pointer {
	p := mmap(v, n, prot, flags, fd, offset)
	// On some systems, mmap ignores v without
	// MAP_FIXED, so retry if the address space is free.
	if p != v && addrspace_free(v, n) {
		if uintptr(p) > 4096 {
			munmap(p, n)
		}
		p = mmap(v, n, prot, flags|MAP_FIXED, fd, offset)
	}
	return p
}

// Don't split the stack as this method may be invoked without a valid G, which
// prevents us from allocating more stack.
//go:nosplit
func sysAlloc(n uintptr, sysStat *uint64) unsafe.Pointer {
	p := mmap(nil, n, PROT_READ|PROT_WRITE, MAP_ANON|MAP_PRIVATE, -1, 0)
	if uintptr(p) < 4096 {
		if uintptr(p) == EACCES {
			print("runtime: mmap: access denied\n")
			exit(2)
		}
		if uintptr(p) == EAGAIN {
			print("runtime: mmap: too much locked memory (check 'ulimit -l').\n")
			exit(2)
		}
		return nil
	}
	mSysStatInc(sysStat, n)
	return p
}

func sysUnused(v unsafe.Pointer, n uintptr) {
	/* madvise(v, n, _MADV_DONTNEED) */
}

func sysUsed(v unsafe.Pointer, n uintptr) {
}

// Don't split the stack as this function may be invoked without a valid G,
// which prevents us from allocating more stack.
//go:nosplit
func sysFree(v unsafe.Pointer, n uintptr, sysStat *uint64) {
	mSysStatDec(sysStat, n)
	munmap(v, n)
}

func sysFault(v unsafe.Pointer, n uintptr) {
	mmap(v, n, PROT_NONE, MAP_ANON|MAP_PRIVATE|MAP_FIXED, -1, 0)
}

func sysReserve(v unsafe.Pointer, n uintptr, reserved *bool) unsafe.Pointer {
	// On 64-bit, people with ulimit -v set complain if we reserve too
	// much address space.  Instead, assume that the reservation is okay
	// if we can reserve at least 64K and check the assumption in SysMap.
	// Only user-mode Linux (UML) rejects these requests.
	if ptrSize == 8 && uint64(n) > 1<<32 {
		p := mmap_fixed(v, 64<<10, PROT_NONE, MAP_ANON|MAP_PRIVATE, -1, 0)
		if p != v {
			if uintptr(p) >= 4096 {
				munmap(p, 64<<10)
			}
			return nil
		}
		munmap(p, 64<<10)
		*reserved = false
		return v
	}

	p := mmap(v, n, PROT_NONE, MAP_ANON|MAP_PRIVATE, -1, 0)
	if uintptr(p) < 4096 {
		return nil
	}
	*reserved = true
	return p
}

func sysMap(v unsafe.Pointer, n uintptr, reserved bool, sysStat *uint64) {
	mSysStatInc(sysStat, n)

	// On 64-bit, we don't actually have v reserved, so tread carefully.
	if !reserved {
		p := mmap_fixed(v, n, PROT_READ|PROT_WRITE, MAP_ANON|MAP_PRIVATE, -1, 0)
		if uintptr(p) == ENOMEM {
			throw("runtime: out of memory")
		}
		if p != v {
			print("runtime: address space conflict: map(", v, ") = ", p, "\n")
			throw("runtime: address space conflict")
		}
		return
	}

	p := mmap(v, n, PROT_READ|PROT_WRITE, MAP_ANON|MAP_FIXED|MAP_PRIVATE, -1, 0)
	if uintptr(p) == ENOMEM {
		throw("runtime: out of memory")
	}
	if p != v {
		throw("runtime: cannot map pages in arena address space")
	}
}


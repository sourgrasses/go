// mksyscall_haiku.pl -tags haiku,amd64 syscall_haiku.go syscall_haiku_amd64.go
// Code generated by the command above; DO NOT EDIT.

//go:build haiku && amd64

package syscall

import "unsafe"

//go:cgo_import_dynamic libc_getgroups getgroups "libroot.so"
//go:cgo_import_dynamic libc_setgroups setgroups "libroot.so"
//go:cgo_import_dynamic libc_fcntl fcntl "libroot.so"
//go:cgo_import_dynamic libc_accept accept "libnetwork.so"
//go:cgo_import_dynamic libc_sendmsg sendmsg "libnetwork.so"
//go:cgo_import_dynamic libc_Access access "libroot.so"
//go:cgo_import_dynamic libc_Adjtime adjtime "libroot.so"
//go:cgo_import_dynamic libc_Chdir chdir "libroot.so"
//go:cgo_import_dynamic libc_Chmod chmod "libroot.so"
//go:cgo_import_dynamic libc_Chown chown "libroot.so"
//go:cgo_import_dynamic libc_Chroot chroot "libroot.so"
//go:cgo_import_dynamic libc_Close close "libroot.so"
//go:cgo_import_dynamic libc_Closedir closedir "libroot.so"
//go:cgo_import_dynamic libc_Dup dup "libroot.so"
//go:cgo_import_dynamic libc_Dup2 dup2 "libroot.so"
//go:cgo_import_dynamic libc_Exit exit "libroot.so"
//go:cgo_import_dynamic libc_Fchdir fchdir "libroot.so"
//go:cgo_import_dynamic libc_Fchmod fchmod "libroot.so"
//go:cgo_import_dynamic libc_Fchown fchown "libroot.so"
//go:cgo_import_dynamic libc_Fdopendir fdopendir "libroot.so"
//go:cgo_import_dynamic libc_Fpathconf fpathconf "libroot.so"
//go:cgo_import_dynamic libc_Fstat fstat "libroot.so"
//go:cgo_import_dynamic libc_Getgid getgid "libroot.so"
//go:cgo_import_dynamic libc_Getpid getpid "libroot.so"
//go:cgo_import_dynamic libc_Geteuid geteuid "libroot.so"
//go:cgo_import_dynamic libc_Getegid getegid "libroot.so"
//go:cgo_import_dynamic libc_Getppid getppid "libroot.so"
//go:cgo_import_dynamic libc_Getpriority getpriority "libroot.so"
//go:cgo_import_dynamic libc_Getrlimit getrlimit "libroot.so"
//go:cgo_import_dynamic libc_Gettimeofday gettimeofday "libroot.so"
//go:cgo_import_dynamic libc_Getuid getuid "libroot.so"
//go:cgo_import_dynamic libc_Kill kill "libroot.so"
//go:cgo_import_dynamic libc_Lchown lchown "libroot.so"
//go:cgo_import_dynamic libc_Link link "libroot.so"
//go:cgo_import_dynamic libc_listen listen "libnetwork.so"
//go:cgo_import_dynamic libc_Lstat lstat "libroot.so"
//go:cgo_import_dynamic libc_Mkdir mkdir "libroot.so"
//go:cgo_import_dynamic libc_Mknod mknod "libroot.so"
//go:cgo_import_dynamic libc_Nanosleep nanosleep "libroot.so"
//go:cgo_import_dynamic libc_Open open "libroot.so"
//go:cgo_import_dynamic libc_Openat openat "libroot.so"
//go:cgo_import_dynamic libc_Pathconf pathconf "libroot.so"
//go:cgo_import_dynamic libc_Pread pread "libroot.so"
//go:cgo_import_dynamic libc_Pwrite pwrite "libroot.so"
//go:cgo_import_dynamic libc_read read "libroot.so"
//go:cgo_import_dynamic libc_Readdir_r readdir_r "libroot.so"
//go:cgo_import_dynamic libc_Readlink readlink "libroot.so"
//go:cgo_import_dynamic libc_Rename rename "libroot.so"
//go:cgo_import_dynamic libc_Rmdir rmdir "libroot.so"
//go:cgo_import_dynamic libc_lseek lseek "libroot.so"
//go:cgo_import_dynamic libc_Setegid setegid "libroot.so"
//go:cgo_import_dynamic libc_Seteuid seteuid "libroot.so"
//go:cgo_import_dynamic libc_Setgid setgid "libroot.so"
//go:cgo_import_dynamic libc_Setpgid setpgid "libroot.so"
//go:cgo_import_dynamic libc_Setpriority setpriority "libroot.so"
//go:cgo_import_dynamic libc_Setregid setregid "libroot.so"
//go:cgo_import_dynamic libc_Setreuid setreuid "libroot.so"
//go:cgo_import_dynamic libc_Setrlimit setrlimit "libroot.so"
//go:cgo_import_dynamic libc_Setsid setsid "libroot.so"
//go:cgo_import_dynamic libc_Setuid setuid "libroot.so"
//go:cgo_import_dynamic libc_shutdown shutdown "libnetwork.so"
//go:cgo_import_dynamic libc_Stat stat "libroot.so"
//go:cgo_import_dynamic libc_Symlink symlink "libroot.so"
//go:cgo_import_dynamic libc_Sync sync "libroot.so"
//go:cgo_import_dynamic libc_Truncate truncate "libroot.so"
//go:cgo_import_dynamic libc_Fsync fsync "libroot.so"
//go:cgo_import_dynamic libc_Ftruncate ftruncate "libroot.so"
//go:cgo_import_dynamic libc_Umask umask "libroot.so"
//go:cgo_import_dynamic libc_Unlink unlink "libroot.so"
//go:cgo_import_dynamic libc_Utimes utimes "libroot.so"
//go:cgo_import_dynamic libc_bind bind "libnetwork.so"
//go:cgo_import_dynamic libc_connect connect "libnetwork.so"
//go:cgo_import_dynamic libc_mmap mmap "libroot.so"
//go:cgo_import_dynamic libc_munmap munmap "libroot.so"
//go:cgo_import_dynamic libc_sendto sendto "libnetwork.so"
//go:cgo_import_dynamic libc_socket socket "libnetwork.so"
//go:cgo_import_dynamic libc_socketpair socketpair "libnetwork.so"
//go:cgo_import_dynamic libc_write write "libroot.so"
//go:cgo_import_dynamic libc_getsockopt getsockopt "libnetwork.so"
//go:cgo_import_dynamic libc_getpeername getpeername "libnetwork.so"
//go:cgo_import_dynamic libc_getsockname getsockname "libnetwork.so"
//go:cgo_import_dynamic libc_setsockopt setsockopt "libnetwork.so"
//go:cgo_import_dynamic libc_recvfrom recvfrom "libnetwork.so"
//go:cgo_import_dynamic libc_recvmsg recvmsg "libnetwork.so"

//go:linkname libc_getgroups libc_getgroups
//go:linkname libc_setgroups libc_setgroups
//go:linkname libc_fcntl libc_fcntl
//go:linkname libc_accept libc_accept
//go:linkname libc_sendmsg libc_sendmsg
//go:linkname libc_Access libc_Access
//go:linkname libc_Adjtime libc_Adjtime
//go:linkname libc_Chdir libc_Chdir
//go:linkname libc_Chmod libc_Chmod
//go:linkname libc_Chown libc_Chown
//go:linkname libc_Chroot libc_Chroot
//go:linkname libc_Close libc_Close
//go:linkname libc_Closedir libc_Closedir
//go:linkname libc_Dup libc_Dup
//go:linkname libc_Dup2 libc_Dup2
//go:linkname libc_Exit libc_Exit
//go:linkname libc_Fchdir libc_Fchdir
//go:linkname libc_Fchmod libc_Fchmod
//go:linkname libc_Fchown libc_Fchown
//go:linkname libc_Fdopendir libc_Fdopendir
//go:linkname libc_Fpathconf libc_Fpathconf
//go:linkname libc_Fstat libc_Fstat
//go:linkname libc_Getgid libc_Getgid
//go:linkname libc_Getpid libc_Getpid
//go:linkname libc_Geteuid libc_Geteuid
//go:linkname libc_Getegid libc_Getegid
//go:linkname libc_Getppid libc_Getppid
//go:linkname libc_Getpriority libc_Getpriority
//go:linkname libc_Getrlimit libc_Getrlimit
//go:linkname libc_Gettimeofday libc_Gettimeofday
//go:linkname libc_Getuid libc_Getuid
//go:linkname libc_Kill libc_Kill
//go:linkname libc_Lchown libc_Lchown
//go:linkname libc_Link libc_Link
//go:linkname libc_listen libc_listen
//go:linkname libc_Lstat libc_Lstat
//go:linkname libc_Mkdir libc_Mkdir
//go:linkname libc_Mknod libc_Mknod
//go:linkname libc_Nanosleep libc_Nanosleep
//go:linkname libc_Open libc_Open
//go:linkname libc_Openat libc_Openat
//go:linkname libc_Pathconf libc_Pathconf
//go:linkname libc_Pread libc_Pread
//go:linkname libc_Pwrite libc_Pwrite
//go:linkname libc_read libc_read
//go:linkname libc_Readdir_r libc_Readdir_r
//go:linkname libc_Readlink libc_Readlink
//go:linkname libc_Rename libc_Rename
//go:linkname libc_Rmdir libc_Rmdir
//go:linkname libc_lseek libc_lseek
//go:linkname libc_Setegid libc_Setegid
//go:linkname libc_Seteuid libc_Seteuid
//go:linkname libc_Setgid libc_Setgid
//go:linkname libc_Setpgid libc_Setpgid
//go:linkname libc_Setpriority libc_Setpriority
//go:linkname libc_Setregid libc_Setregid
//go:linkname libc_Setreuid libc_Setreuid
//go:linkname libc_Setrlimit libc_Setrlimit
//go:linkname libc_Setsid libc_Setsid
//go:linkname libc_Setuid libc_Setuid
//go:linkname libc_shutdown libc_shutdown
//go:linkname libc_Stat libc_Stat
//go:linkname libc_Symlink libc_Symlink
//go:linkname libc_Sync libc_Sync
//go:linkname libc_Truncate libc_Truncate
//go:linkname libc_Fsync libc_Fsync
//go:linkname libc_Ftruncate libc_Ftruncate
//go:linkname libc_Umask libc_Umask
//go:linkname libc_Unlink libc_Unlink
//go:linkname libc_Utimes libc_Utimes
//go:linkname libc_bind libc_bind
//go:linkname libc_connect libc_connect
//go:linkname libc_mmap libc_mmap
//go:linkname libc_munmap libc_munmap
//go:linkname libc_sendto libc_sendto
//go:linkname libc_socket libc_socket
//go:linkname libc_socketpair libc_socketpair
//go:linkname libc_write libc_write
//go:linkname libc_getsockopt libc_getsockopt
//go:linkname libc_getpeername libc_getpeername
//go:linkname libc_getsockname libc_getsockname
//go:linkname libc_setsockopt libc_setsockopt
//go:linkname libc_recvfrom libc_recvfrom
//go:linkname libc_recvmsg libc_recvmsg

type libcFunc uintptr

var (
	libc_getgroups,
	libc_setgroups,
	libc_fcntl,
	libc_accept,
	libc_sendmsg,
	libc_Access,
	libc_Adjtime,
	libc_Chdir,
	libc_Chmod,
	libc_Chown,
	libc_Chroot,
	libc_Close,
	libc_Closedir,
	libc_Dup,
	libc_Dup2,
	libc_Exit,
	libc_Fchdir,
	libc_Fchmod,
	libc_Fchown,
	libc_Fdopendir,
	libc_Fpathconf,
	libc_Fstat,
	libc_Getgid,
	libc_Getpid,
	libc_Geteuid,
	libc_Getegid,
	libc_Getppid,
	libc_Getpriority,
	libc_Getrlimit,
	libc_Gettimeofday,
	libc_Getuid,
	libc_Kill,
	libc_Lchown,
	libc_Link,
	libc_listen,
	libc_Lstat,
	libc_Mkdir,
	libc_Mknod,
	libc_Nanosleep,
	libc_Open,
	libc_Openat,
	libc_Pathconf,
	libc_Pread,
	libc_Pwrite,
	libc_read,
	libc_Readdir_r,
	libc_Readlink,
	libc_Rename,
	libc_Rmdir,
	libc_lseek,
	libc_Setegid,
	libc_Seteuid,
	libc_Setgid,
	libc_Setpgid,
	libc_Setpriority,
	libc_Setregid,
	libc_Setreuid,
	libc_Setrlimit,
	libc_Setsid,
	libc_Setuid,
	libc_shutdown,
	libc_Stat,
	libc_Symlink,
	libc_Sync,
	libc_Truncate,
	libc_Fsync,
	libc_Ftruncate,
	libc_Umask,
	libc_Unlink,
	libc_Utimes,
	libc_bind,
	libc_connect,
	libc_mmap,
	libc_munmap,
	libc_sendto,
	libc_socket,
	libc_socketpair,
	libc_write,
	libc_getsockopt,
	libc_getpeername,
	libc_getsockname,
	libc_setsockopt,
	libc_recvfrom,
	libc_recvmsg libcFunc
)

func getgroups(ngid int, gid *_Gid_t) (n int, err error) {
	r0, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_getgroups)), 2, uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0, 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func setgroups(ngid int, gid *_Gid_t) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_setgroups)), 2, uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func fcntl(fd int, cmd int, arg int) (val int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_fcntl)), 3, uintptr(fd), uintptr(cmd), uintptr(arg), 0, 0, 0)
	val = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_accept)), 3, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func sendmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_sendmsg)), 3, uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Access(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Access)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Adjtime(delta *Timeval, olddelta *Timeval) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Adjtime)), 2, uintptr(unsafe.Pointer(delta)), uintptr(unsafe.Pointer(olddelta)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Chdir(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Chdir)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Chmod(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Chmod)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Chown(path string, uid int, gid int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Chown)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Chroot(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Chroot)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Close(fd int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Close)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Closedir(dir uintptr) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Closedir)), 1, uintptr(dir), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Dup(fd int) (nfd int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Dup)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	nfd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Dup2(from int, to int) (nfd int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Dup2)), 2, uintptr(from), uintptr(to), 0, 0, 0, 0)
	nfd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Fchdir(fd int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Fchdir)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Fchmod(fd int, mode uint32) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Fchmod)), 2, uintptr(fd), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Fchown(fd int, uid int, gid int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Fchown)), 3, uintptr(fd), uintptr(uid), uintptr(gid), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Fdopendir(fd int) (dir uintptr, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Fdopendir)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	dir = uintptr(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Fpathconf(fd int, name int) (val int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Fpathconf)), 2, uintptr(fd), uintptr(name), 0, 0, 0, 0)
	val = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Fstat(fd int, stat *Stat_t) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Fstat)), 2, uintptr(fd), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Getgid() (gid int) {
	r0, _, _ := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Getgid)), 0, 0, 0, 0, 0, 0, 0)
	gid = int(r0)
	return
}

func Getpid() (pid int) {
	r0, _, _ := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Getpid)), 0, 0, 0, 0, 0, 0, 0)
	pid = int(r0)
	return
}

func Geteuid() (euid int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&libc_Geteuid)), 0, 0, 0, 0, 0, 0, 0)
	euid = int(r0)
	return
}

func Getegid() (egid int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&libc_Getegid)), 0, 0, 0, 0, 0, 0, 0)
	egid = int(r0)
	return
}

func Getppid() (ppid int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&libc_Getppid)), 0, 0, 0, 0, 0, 0, 0)
	ppid = int(r0)
	return
}

func Getpriority(which int, who int) (n int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Getpriority)), 2, uintptr(which), uintptr(who), 0, 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Getrlimit(which int, lim *Rlimit) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Getrlimit)), 2, uintptr(which), uintptr(unsafe.Pointer(lim)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Gettimeofday(tv *Timeval) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Gettimeofday)), 1, uintptr(unsafe.Pointer(tv)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Getuid() (uid int) {
	r0, _, _ := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Getuid)), 0, 0, 0, 0, 0, 0, 0)
	uid = int(r0)
	return
}

func Kill(pid int, signum Signal) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Kill)), 2, uintptr(pid), uintptr(signum), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Lchown(path string, uid int, gid int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Lchown)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Link(path string, link string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(link)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Link)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Listen(s int, backlog int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_listen)), 2, uintptr(s), uintptr(backlog), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Lstat(path string, stat *Stat_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Lstat)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Mkdir(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Mkdir)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Mknod(path string, mode uint32, dev int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Mknod)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(dev), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Nanosleep(time *Timespec, leftover *Timespec) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Nanosleep)), 2, uintptr(unsafe.Pointer(time)), uintptr(unsafe.Pointer(leftover)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Open(path string, mode int, perm uint32) (fd int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Open)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(perm), 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Openat(fd int, path string, flags int, perm uint32) (fdret int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Openat)), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(flags), uintptr(perm), 0, 0)
	fdret = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Pathconf(path string, name int) (val int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Pathconf)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(name), 0, 0, 0, 0)
	val = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Pread(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Pread)), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Pwrite(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Pwrite)), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func read(fd int, p []byte) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_read)), 3, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Readdir_r(dir uintptr, entry *Dirent, result **Dirent) (res Errno) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&libc_Readdir_r)), 3, uintptr(dir), uintptr(unsafe.Pointer(entry)), uintptr(unsafe.Pointer(result)), 0, 0, 0)
	res = Errno(r0)
	return
}

func Readlink(path string, buf []byte) (n int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	if len(buf) > 0 {
		_p1 = &buf[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Readlink)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), uintptr(len(buf)), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Rename(from string, to string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(from)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(to)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Rename)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Rmdir(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Rmdir)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Seek(fd int, offset int64, whence int) (newoffset int64, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_lseek)), 3, uintptr(fd), uintptr(offset), uintptr(whence), 0, 0, 0)
	newoffset = int64(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Setegid(egid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Setegid)), 1, uintptr(egid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Seteuid(euid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Seteuid)), 1, uintptr(euid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Setgid(gid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Setgid)), 1, uintptr(gid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Setpgid(pid int, pgid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Setpgid)), 2, uintptr(pid), uintptr(pgid), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Setpriority(which int, who int, prio int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Setpriority)), 3, uintptr(which), uintptr(who), uintptr(prio), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Setregid(rgid int, egid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Setregid)), 2, uintptr(rgid), uintptr(egid), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Setreuid(ruid int, euid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Setreuid)), 2, uintptr(ruid), uintptr(euid), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Setrlimit(which int, lim *Rlimit) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Setrlimit)), 2, uintptr(which), uintptr(unsafe.Pointer(lim)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Setsid() (pid int, err error) {
	r0, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Setsid)), 0, 0, 0, 0, 0, 0, 0)
	pid = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Setuid(uid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_Setuid)), 1, uintptr(uid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Shutdown(s int, how int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_shutdown)), 2, uintptr(s), uintptr(how), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Stat(path string, stat *Stat_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Stat)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Symlink(path string, link string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(link)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Symlink)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Sync() (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Sync)), 0, 0, 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Truncate(path string, length int64) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Truncate)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(length), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Fsync(fd int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Fsync)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Ftruncate(fd int, length int64) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Ftruncate)), 2, uintptr(fd), uintptr(length), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Umask(newmask int) (oldmask int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&libc_Umask)), 1, uintptr(newmask), 0, 0, 0, 0, 0)
	oldmask = int(r0)
	return
}

func Unlink(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Unlink)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func Utimes(path string, times *[2]Timeval) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_Utimes)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(times)), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_bind)), 3, uintptr(s), uintptr(addr), uintptr(addrlen), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_connect)), 3, uintptr(s), uintptr(addr), uintptr(addrlen), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func mmap(addr uintptr, length uintptr, prot int, flag int, fd int, pos int64) (ret uintptr, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_mmap)), 6, uintptr(addr), uintptr(length), uintptr(prot), uintptr(flag), uintptr(fd), uintptr(pos))
	ret = uintptr(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func munmap(addr uintptr, length uintptr) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_munmap)), 2, uintptr(addr), uintptr(length), 0, 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_sendto)), 6, uintptr(s), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(flags), uintptr(to), uintptr(addrlen))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func socket(domain int, typ int, proto int) (fd int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_socket)), 3, uintptr(domain), uintptr(typ), uintptr(proto), 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func socketpair(domain int, typ int, proto int, fd *[2]int32) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_socketpair)), 4, uintptr(domain), uintptr(typ), uintptr(proto), uintptr(unsafe.Pointer(fd)), 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func write(fd int, p []byte) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_write)), 3, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_getsockopt)), 5, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(unsafe.Pointer(vallen)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&libc_getpeername)), 3, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_getsockname)), 3, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_setsockopt)), 5, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(vallen), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_recvfrom)), 6, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(flags), uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(fromlen)))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func recvmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&libc_recvmsg)), 3, uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || (solaris && !illumos)

package filelock

type inode = uint64 // type of syscall.Stat_t.Ino

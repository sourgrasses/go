// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"syscall"
	"time"
)

func sameFile(fs1, fs2 *fileStat) bool {
	//STUB
	panic("Not implemented")
}

func fileInfoFromStat(st *syscall.Stat_t, name string) FileInfo {
	panic("Not implemented")
}

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

// For testing.
func atime(fi FileInfo) time.Time {
	panic("Not implemented")
}

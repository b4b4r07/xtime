package xtime

import (
	"os"
	"syscall"
	"time"
)

func ctime(fi os.FileInfo) time.Time {
	return time.Unix(0, fi.Sys().(*syscall.Win32FileAttributeData).LastAccessTime.Nanoseconds())
}

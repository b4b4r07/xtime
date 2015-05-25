package xtime

import (
	"os"
	"time"
)

func atime(fi os.FileInfo) time.Time {
	return time.Unix(int64(fi.Sys().(*syscall.Dir).Atime), 0)
}

func ctime(fi os.FileInfo) time.Time {
	return time.Unix(int64(fi.Sys().(*syscall.Dir).Ctime), 0)
}

func mtime(fi os.FileInfo) time.Time {
	return time.Unix(int64(fi.Sys().(*syscall.Dir).Mtime), 0)
}

// Package ctime provides a platform-independent way to get ctimes for files.
package xtime

import (
	"os"
	"time"
)

func Atime(fi os.FileInfo) time.Time {
	return atime(fi)
}

func Ctime(fi os.FileInfo) time.Time {
	return ctime(fi)
}

func Mtime(fi os.FileInfo) time.Time {
	return mtime(fi)
}

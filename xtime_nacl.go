package xtime

import (
	"os"
	"syscall"
	"time"
)

func timespecToTime(sec, nsec int64) time.Time {
	return time.Unix(sec, nsec)
}

func atime(fi os.FileInfo) time.Time {
	st := fi.Sys().(*syscall.Stat_t)
	return timespecToTime(st.Atime, st.AtimeNsec)
}

func ctime(fi os.FileInfo) time.Time {
	st := fi.Sys().(*syscall.Stat_t)
	return timespecToTime(st.Ctime, st.CtimeNsec)
}

func mtime(fi os.FileInfo) time.Time {
	st := fi.Sys().(*syscall.Stat_t)
	return timespecToTime(st.Mtime, st.MtimeNsec)
}

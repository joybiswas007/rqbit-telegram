package utils

import (
	"fmt"
	"runtime"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

// DiskUsage returns the total, free, and used disk space in bytes.
func DiskUsage() (total uint64, free uint64, used uint64, err error) {
	var stat syscall.Statfs_t
	err = syscall.Statfs("/", &stat)
	if err != nil {
		return
	}
	total = stat.Blocks * uint64(stat.Bsize)
	free = stat.Bfree * uint64(stat.Bsize)
	used = total - free
	return
}

// MemoryUsage returns the total, free, and used memory in bytes.
func MemoryUsage() (total uint64, free uint64, used uint64) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	total = memStats.Sys
	free = memStats.HeapIdle
	used = total - free
	return
}

// MemoryUsage returns the total, free, and used memory in bytes.
func RAMUsage() (total uint64, free uint64, used uint64) {
	var info unix.Sysinfo_t
	err := unix.Sysinfo(&info)
	if err != nil {
		panic(fmt.Sprintf("Failed to get system memory info: %v", err))
	}
	total = info.Totalram * uint64(info.Unit)
	free = info.Freeram * uint64(info.Unit)
	used = total - free
	return
}

// Uptime returns the duration for which the system has been up.
func Uptime() (uptime time.Duration, err error) {
	var info syscall.Sysinfo_t
	err = syscall.Sysinfo(&info)
	if err != nil {
		return 0, err
	}
	uptime = time.Duration(info.Uptime) * time.Second
	return
}

// HumanReadableBytes converts bytes to a human-readable string.
func HumanReadableBytes(bytes uint64) string {
	const (
		_         = iota
		kb uint64 = 1 << (10 * iota)
		mb
		gb
		tb
		pb
	)

	switch {
	case bytes >= pb:
		return fmt.Sprintf("%.2f PB", float64(bytes)/float64(pb))
	case bytes >= tb:
		return fmt.Sprintf("%.2f TB", float64(bytes)/float64(tb))
	case bytes >= gb:
		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(gb))
	case bytes >= mb:
		return fmt.Sprintf("%.2f MB", float64(bytes)/float64(mb))
	case bytes >= kb:
		return fmt.Sprintf("%.2f KB", float64(bytes)/float64(kb))
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}

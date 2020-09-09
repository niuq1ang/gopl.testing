package main

import (
	"fmt"
	"syscall"
)

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

// DiskUsage disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	fmt.Println("all: ", disk.All/(1024*1024*1024))
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	fmt.Println("free: ", disk.Free/(1024*1024*1024))
	disk.Used = disk.All - disk.Free
	fmt.Println("used: ", disk.Used/(1024*1024*1024))
	return
}

func main() {

	fmt.Printf("%+v\n", DiskUsage("/"))
}

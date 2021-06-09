// +build !windows
package gcdu

import (
	"fmt"

	syscall "golang.org/x/sys/unix"
)

const (
	//Byte conversion factors
	B_MB = (1024.0 * 1024)
	B_GB = (1024.0 * 1024.0 * 1024.0)
)

type UsageInfo struct {
	Total uint64
	Used  uint64
	Free  uint64
}

func FetchDiskUsage(path string) (disk UsageInfo) {

	// defining Statfs struct
	var fs_t syscall.Statfs_t = syscall.Statfs_t{}
	// call Statfs and returns the status
	err := syscall.Statfs(path, &fs_t)
	if err != nil {
		fmt.Printf("[ERROR]: %s\n", err)
		return
	} else {

		//updating disk info
		disk.Total = fs_t.Blocks * uint64(fs_t.Bsize)
		disk.Free = fs_t.Bfree * uint64(fs_t.Bsize)
		disk.Used = disk.Total - disk.Free
		return
	}
}

func DisplayInfo(disk UsageInfo) {

	// Display usage info:
	fmt.Println("===============\nDisk Usage Info\n===============")
	fmt.Printf("Total\t:\t%.1f GB\n",
		float64(disk.Total)/B_GB)
	fmt.Printf("Used\t:\t%.1f GB\n",
		float64(disk.Used)/B_GB)
	fmt.Printf("Free\t:\t%.1f GB\n",
		float64(disk.Free)/B_GB)

}

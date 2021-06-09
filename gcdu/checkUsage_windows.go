// +build windows
package gcdu

import (
	"log"
	"syscall"
	"unsafe"
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
	// Loading module that handles memory usage of Windows:
	kernel, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		log.Panic(err)
	}
	defer syscall.FreeLibrary(kernel)

	// Retrieve process address:
	GetDiskFreeSpaceEx, err := syscall.GetProcAddress(syscall.Handle(kernel), "GetDiskFreeSpaceExW")
	if err != nil {
		log.Panic(err)

	}

	syscall.Syscall6(uintptr(GetDiskFreeSpaceEx), 4,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(path))), 0,
		uintptr(unsafe.Pointer(&disk.Total)),
		uintptr(unsafe.Pointer(&disk.Free)), 0, 0)
	disk.Used = disk.Total - disk.Free
	return
}

func DisplayInfo(disk UsageInfo) {

	log.Println("===============\nDisk Usage Info\n===============")
	log.Printf("Total\t:\t%.1f GB", float64(disk.Total)/B_GB)
	log.Printf("Used\t:\t%.1f GB", float64(disk.Used)/B_GB)
	log.Printf("Free\t:\t%.1f GB", float64(disk.Free)/B_GB)

}

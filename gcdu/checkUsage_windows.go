// +build windows
package gcdu

import (
	"flag"
	"log"
	"syscall"
	"unsafe"
)

const (
	//Byte conversion factors
	B_MB = (1024.0 * 1024)
	B_GB = (1024.0 * 1024.0 * 1024.0)
)

type DiskInfo struct {
	Total uint64
	Used  uint64
	Free  uint64
}

func getDiskInfo(path string) (disk DiskInfo) {
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

func displayInfo(path string) {
	// obtain disk usage info
	info := getDiskInfo(path)
	if info.Total == 0 {
		log.Panic("[ERROR]: Invalid path")
		return
	}

	log.Println("===============\nDisk Usage Info\n===============")
	log.Printf("Total\t:\t%.1f GB", float64(info.Total)/B_GB)
	log.Printf("Used\t:\t%.1f GB", float64(info.Used)/B_GB)
	log.Printf("Free\t:\t%.1f GB", float64(info.Free)/B_GB)

}

func main() {

	path := flag.String("path", "", "A valid file path")
	flag.Parse()

	displayInfo(*path)

}

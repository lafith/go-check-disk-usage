# go-check-disk-usage

Check disk usage information like total space, used and free.
Implemented in Go using syscall.

## Demo:
```go
package main

import ("github.com/lafith/go-check-disk-usage/gcdu")

func main() {
	info := gcdu.FetchDiskUsage("/home")
	gcdu.DisplayInfo(info)
}
```

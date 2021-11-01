/**package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Reciving")
	handle, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, 0x1)
	f := os.NewFile(uintptr(handle), fmt.Sprintf("fd %d", handle))
	for {
		buf := make([]byte, 1024)
		numRead, err := f.Read(buf)
		if err != nil {
			fmt.Println("OOOOO")
			fmt.Println(err)
		}
		fmt.Println(numRead)
		//fmt.Printf(" X\n", buf[:numRead])
		time.Sleep(time.Second)
	}
}**/

package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fd, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, 0x1)
	f := os.NewFile(uintptr(fd), fmt.Sprintf("fd %d", fd))

	for {
		buf := make([]byte, 1024)
		numRead, err := f.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("% X\n", buf[:numRead])
	}
}

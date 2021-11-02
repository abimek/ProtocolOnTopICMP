package main

import (
	"fmt"
	"golang.org/x/sys/unix"
)

func main() {
	fd, error := unix.Socket(unix.AF_INET, unix.SOCK_RAW, unix.IPPROTO_ICMP)
	if error != nil {
		unix.Close(fd)
		panic(error)
	}
	data := make([]byte, 1024)
	for {
		unix.Recvfrom(fd, data, 0)
		fmt.Println(data)
	}
}

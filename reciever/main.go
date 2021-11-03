package main

import (
	"encoding/binary"
	"fmt"
	"golang.org/x/sys/unix"
)

type ipheader struct {
	VL             uint8
	tos            uint8
	totallength    uint16
	identification uint16
	ffo            uint16
	ttl            uint8
	protocol       uint8
	checksum       uint16
	srcAddr        [4]byte
	dstAddr        [4]byte
}

type CPHeader struct {
	//srcprt uint16
	//dstprt uint16
	sndrage uint8
	zipcode uint32
}

func main() {
	fd, error := unix.Socket(unix.AF_INET, unix.SOCK_RAW, unix.IPPROTO_RAW)
	if error != nil {
		unix.Close(fd)
		panic(error)
	}
	data := make([]byte, 1024)
	unix.Recvfrom(fd, data, 0)
	iphdr, cphdr := readData(data)
	fmt.Printf("%+v\n", iphdr)
	fmt.Println("--------=")
	fmt.Printf("%+v\n", cphdr)
}

func readData(data []byte) (ipheader, CPHeader) {
	IPHDR := ipheader{
		VL:             data[0],
		tos:            data[1],
		totallength:    binary.BigEndian.Uint16([]byte{data[2], data[3]}),
		identification: binary.BigEndian.Uint16([]byte{data[4], data[5]}),
		ffo:            binary.BigEndian.Uint16([]byte{data[6], data[7]}),
		ttl:            data[8],
		protocol:       data[9],
		checksum:       binary.BigEndian.Uint16([]byte{data[10], data[11]}),
		srcAddr:        [4]byte{data[12], data[13], data[14], data[15]},
		dstAddr:        [4]byte{data[16], data[17], data[18], data[19]},
	}
	CPHDR := CPHeader{
		sndrage: data[30],
		zipcode: binary.BigEndian.Uint32([]byte{data[31], data[32], data[33], data[34]}),
	}

	return IPHDR, CPHDR
}

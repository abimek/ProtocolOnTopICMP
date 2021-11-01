package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"syscall"
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
	fmt.Println("sending")
	//TODO IMPLEMENT READING IP AND SRC FROM FILE
	srcipa := "151.200.239.125"
	srcipb := net.IP{}
	if srcipb = net.ParseIP(srcipa); srcipb == nil {
		fmt.Printf("IP Address: %s - Invalid\n", srcipa)
	} else {
		fmt.Printf("IP Address: %s - Valid\n", srcipa)
	}
	dstripa := "151.200.239.125"
	dstipb := net.IP{}
	if dstipb = net.ParseIP(dstripa); dstipb == nil {
		fmt.Printf("IP Address: %s - Invalid\n", dstripa)
	} else {
		fmt.Printf("IP Address: %s - Valid\n", dstripa)
	}
	srcip := [4]byte{}
	dstip := [4]byte{}
	copy(srcip[:], srcipb)
	copy(dstip[:], dstipb)
	//oxff is IPPROTO_RAW
	handle, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, 0)
	if err != nil {
		fmt.Println(err)
		fmt.Println("err1")
	}

	err2 := syscall.SetsockoptInt(handle, syscall.IPPROTO_IP, 0x3, 1)
	if err2 != nil {
		fmt.Println(err)
		fmt.Println("err2Bitrch")
	}
	address := syscall.SockaddrInet4{
		Port: 0,
		Addr: dstip,
	}

	//CREATE PAYLOAD

	p := createPayLoad(srcip, dstip)

	errs := syscall.Sendto(handle, p, 0, &address)
	if errs != nil {
		fmt.Println(errs)
		fmt.Println("ERRRROROROROOROR")
	}
	for {

	}
}

func createPayLoad(srcip [4]byte, dstip [4]byte) []byte {
	data := bytes.Buffer{}

	iphdr := ipheader{
		VL:             0x45,
		tos:            0x00,
		identification: 0x1214,
		ffo:            0x0000,
		ttl:            100,
		protocol:       1,
		srcAddr:        srcip,
		dstAddr:        dstip,
	}
	icmp := []byte{
		8,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0xC0,
		0xDE,
	}

	cs := csum(icmp)
	icmp[2] = byte(cs)
	icmp[3] = byte(cs >> 8)

	extraData := CPHeader{
		sndrage: 12,
		zipcode: 20105,
	}

	err := binary.Write(&data, binary.BigEndian, &iphdr)
	if err != nil {
		fmt.Println(err)
	}
	err1 := binary.Write(&data, binary.BigEndian, &icmp)
	if err1 != nil {
		fmt.Println(err1)
	}
	err2 := binary.Write(&data, binary.BigEndian, &extraData)
	if err2 != nil {
		fmt.Println(err2)
	}
	return data.Bytes()
}

//This Portion of Code is barrowed and WIll be rewritten Later
func csum(b []byte) uint16 {
	var s uint32
	for i := 0; i < len(b); i += 2 {
		s += uint32(b[i+1])<<8 | uint32(b[i])
	}
	// add back the carry
	s = s>>16 + s&0xffff
	s = s + s>>16
	return uint16(^s)
}

package main

import (
	"fmt"
	"net"
)

func main() {
	con, err := net.Dial("tcp", "0.0.0.0:5555")
	fmt.Println("Connected to server")
	if err != nil {
		panic(err)
	}
	con.Write([]byte{0, 0, 0, 1})
	buf := make([]byte, 1024)
	for {
		n, err := con.Read(buf[:cap(buf)])
		if err != nil {
			panic(err)
		}

		buf = buf[:n]
		if string(buf) != "ping" {
			fmt.Println("Buffer:", buf)
			length := int(buf[0])<<8 + int(buf[1])
			time := buf[5]
			targets := buf[7]
			ip := net.IPv4(buf[8], buf[9], buf[10], buf[11])
			flag := buf[14]
			ports := buf[15]
			port := string(buf[16:])
			fmt.Println("Length: ", length)
			fmt.Println("Time: ", time)
			fmt.Println("Targets: ", targets)
			fmt.Println("IP: ", ip)
			fmt.Println("Flag: ", flag)
			fmt.Println("Ports: ", ports)
			fmt.Println("Port: ", port)
		}
	}
}

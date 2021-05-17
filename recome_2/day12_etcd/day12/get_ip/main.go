package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// GetOutBoundIP Get Local 对外 IP
func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Printf("net dial err: %v\n", err)
		return
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]

	return
}

func main() {
	ip, err := GetOutBoundIP()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stdout, ip)
}

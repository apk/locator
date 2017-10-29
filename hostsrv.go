// run-on-change hostsrv.go -- g build hostsrv.go -- ./hostsrv

package main

import (
	"fmt"
	"net"
	"os"
)

func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

func main() {
	ServerAddr,err := net.ResolveUDPAddr("udp",":11889")
	CheckError(err)
 
	hostname, _ := os.Hostname()

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()
 
	buf := make([]byte, 1024)
 
	for {
		n,addr,err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received ",string(buf[0:n]), " from ",addr)
 
		if err != nil {
			fmt.Println("Error: ",err)
		} 

		n,err = ServerConn.WriteToUDP([]byte(hostname),addr);
		if err != nil {
			fmt.Println("Send error: ",err)
		} 
	}
}

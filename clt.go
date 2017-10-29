// run-on-change clt.go -- g build clt.go -- ./clt

package main
 
import (
    "fmt"
    "net"
    "time"
    "strconv"
)
 
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}
 
func main() {
	RemAddr,err := net.ResolveUDPAddr("udp","255.255.255.255:11889")
	//RemAddr,err := net.ResolveUDPAddr("udp","192.168.43.255:11889")
	//RemAddr,err := net.ResolveUDPAddr("udp","[ff02::1%wlan0]:11889")
	CheckError(err)
 
	ServerAddr, err := net.ResolveUDPAddr("udp", ":0")
	CheckError(err)
 
	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	//Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	//CheckError(err)
 
	//defer Conn.Close()

	xmit := func () {
		i := 0
		for {
			msg := strconv.Itoa(i)
			i++
			buf := []byte(msg)
			_,err := ServerConn.WriteToUDP(buf,RemAddr)
			if err != nil {
				fmt.Println(msg, err)
			}
			time.Sleep(time.Second * 1)
		}
	}

	go xmit();

	buf := make([]byte, 1024)

	for {
		n,addr,err := ServerConn.ReadFromUDP(buf)
		fmt.Printf("Received %v from %v%%%v\n",string(buf[0:n]), addr.IP,addr.Zone)
 
		if err != nil {
			fmt.Println("Error: ",err)
		} 
	}
}

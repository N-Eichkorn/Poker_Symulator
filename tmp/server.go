/**
* TCPArithServer
 */
package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
	"time"
)

type Element struct {
	Timestamp time.Time
	Hostname  string
	HW_Info   string
}

type List []Element

func (l *List) Add(hostname string, reply *int) error {
	*l = append(*l, Element{Timestamp: time.Now(), Hostname: hostname, HW_Info: ""})

	fmt.Println(l)
	return nil
}

func main() {
	list := new(List)
	rpc.Register(list)
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	/* This works:
	   rpc.Accept(listener)
	*/
	/* and so does this:
	 */
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

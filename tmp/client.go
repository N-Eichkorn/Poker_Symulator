/**
* TCPArithClient
 */
package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"runtime"
)

var (
	command string
)

func checkOS() {
	os := runtime.GOOS
	switch os {
	case "windows":
		command = "systeminfo"
	case "linux":
		command = "lshw"
	default:
		fmt.Printf("%s.\n", os)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server:port")
		os.Exit(1)
	}
	service := os.Args[1]
	client, err := rpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	hostname, _ := os.Hostname()
	var reply int
	err = client.Call("List.Add", hostname, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
}

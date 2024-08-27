package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Demo: tcp client")
	conn, err := net.Dial("tcp", ":1011")
	if err != nil {
		fmt.Println("Received error", err)
		os.Exit(1)
	}
	for {
		n, e := conn.Write([]byte("tcp connection initiated..."))
		if e != nil {
			fmt.Println("Failed to write")
		} else {

			fmt.Println(n, "bytes written..")
			data := make([]byte, 1024)
			n, e = conn.Read(data)
			if e != nil {
				fmt.Println("received error")

			}
			fmt.Println(n, "byets received and received data is", string(data))

		}
	}
}

package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("Demo tcp service")

	list, err := net.Listen("tcp", ":1011")

	if err != nil {
		fmt.Println("Failed to create service")
		os.Exit(1)
	}
	fmt.Println("Service created")
	for {
		conn, err := list.Accept()
		if err != nil {
			fmt.Println("Receive an error")
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	for {
		data := make([]byte, 1024, 1024)
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("Got a read error")
		}
		fmt.Println("Received n number bytes", n)

		n, err = conn.Write([]byte("Connection successfull, will terminate connection"))
		if err != nil {
			fmt.Println("Got a write error")
		}
		time.Sleep(2 * time.Second)
		fmt.Println(n, "bytes written")

		fmt.Println(string(data))
	}
}

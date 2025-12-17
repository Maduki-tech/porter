package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("Server error:", err)
		}
	}()
	fmt.Println("Starting port scan on localhost...")
	for port := 1; port <= 10000; port++ {
		if ScanPort(port) {
			fmt.Printf("Port %d is open\n", port)
		}
	}
}

func ScanPort(port int) bool {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

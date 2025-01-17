package main

import (
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {
	clientpath := filepath.Join(os.TempDir(), "unixdomainsocket-client")
	os.Remove(clientpath)

	conn, err := net.ListenPacket("unixgram", clientpath)
	if err != nil {
		panic(err)
	}

	// 送信先のアドレス
	unixServerAddr, err := net.ResolveUnixAddr("unixgram", filepath.Join(os.TempDir(), "unixdomainsocket-server"))
	var serverAddr net.Addr = unixServerAddr
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	log.Println("Sending to Server")

	_, err = conn.WriteTo([]byte("Hello from Client"), serverAddr)
	if err != nil {
		panic(err)
	}
	log.Println("Receiving from Server")
	buffer := make([]byte, 1500)
	length, _, err := conn.ReadFrom(buffer)
	if err != nil {
		panic(err)
	}
	log.Printf("Received: %s\n", string(buffer[:length]))
}

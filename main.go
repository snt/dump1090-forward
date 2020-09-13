package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	sourceAddr := flag.String("source", "127.0.0.1:30005", "source address host:port")
	targetAddr := flag.String("target", "192.168.10.200:30004", "target address host:port")
	bufferSize := flag.Int("buffer-size", 15384, "buffer size")
	flag.Parse()

	for {
		err := relayLoop(*sourceAddr, *targetAddr, *bufferSize)
		log.Printf("reconectting due to error %v\n", err)
		time.Sleep(1 * time.Second)
	}
}

func relayLoop(sourceAddr string, targetAddr string, bufferSize int) error {
	sourceTcp, err := net.Dial("tcp", sourceAddr)
	if err != nil {
		return fmt.Errorf("failed to connect to source %v", err)
	}
	defer sourceTcp.Close()

	targetTcp, err := net.Dial("tcp", targetAddr)
	if err != nil {
		return fmt.Errorf("failed to connec to target %v", err)
	}
	defer targetTcp.Close()

	log.Printf("connected and relaying from %s to %s", sourceAddr, targetAddr)

	reader := bufio.NewReader(sourceTcp)
	writer := bufio.NewWriter(targetTcp)

	buffer := make([]byte, bufferSize)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			return fmt.Errorf("failed to read %v", err)
		}
		_, err = writer.Write(buffer[:n])
		if err != nil {
			return fmt.Errorf("failed to write %v", err)
		}
	}
}

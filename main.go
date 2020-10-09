package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	sourceAddr := flag.String("source", "127.0.0.1:30005", "source address host:port")
	targetAddr := flag.String("target", "192.168.10.200:30004", "target address host:port")
	flag.Parse()

	for {
		err := relayLoop(*sourceAddr, *targetAddr)
		if err != nil {
			log.Printf("error=[%v]\n", err)
			time.Sleep(1 * time.Second)
		}
		log.Println("reconnecting")
	}
}

func relayLoop(sourceAddr string, targetAddr string) error {
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

	//reader := bufio.NewReader(sourceTcp)
	//writer := bufio.NewWriter(targetTcp)

	wc, err := io.Copy(targetTcp, sourceTcp)
	if err != nil {
		return fmt.Errorf("failed to copy error=[%v]", err)
	}
	log.Printf("EOF received (%d bytes are transfered)", wc)

	return nil
}

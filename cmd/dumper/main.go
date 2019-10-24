package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"time"
)

func main() {

	address := "0.0.0.0:8080"
	l, err := net.Listen("tcp4", address)
	if err != nil {
		log.Fatalf("can't start socker on address %s: %v\n", address, err)
	}
	defer l.Close()

	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}

}

func handleConnection(c net.Conn) {
	defer c.Close()

	fmt.Println()
	reader := bufio.NewReader(c)
	barr, err := reader.ReadBytes('\n')
	if err != nil {
		if err != io.EOF {
			log.Printf("error while reading data: %v\n", err)
		}
		return
	}
	line := string(barr)
	for len(line) > 2 || (line != "\n" && line != "\r\n") {
		fmt.Print(line)
		barr, err = reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("error while reading data: %v\n", err)
			}
			return
		}
		line = string(barr)
	}

	c.Write([]byte("HTTP/1.1 204 No Content\n"))
	c.Write([]byte("Server: Dumper\n"))
	c.Write([]byte("Connection: close\n"))
	c.Write([]byte("\n"))
}

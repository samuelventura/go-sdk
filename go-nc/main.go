package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"time"
)

const NL = '\n'

//go run . 10.77.0.49:31607
func main() {
	addr := os.Args[1]
	conn, err := net.DialTimeout("tcp", addr, 3*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		writer := bufio.NewWriter(os.Stdout)
		reader := bufio.NewReader(conn)
		for {
			line, err := reader.ReadString(NL)
			if err != nil {
				log.Fatal(err)
			}
			writer.WriteString(line)
			writer.Flush()
		}
	}()
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(conn)
	for {
		line, err := reader.ReadString(NL)
		if err != nil {
			log.Fatal(err)
		}
		writer.WriteString(line)
		writer.Flush()
	}
}

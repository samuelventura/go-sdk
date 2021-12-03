package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

const CR = '\r'
const NL = '\n'

//go run . 10.77.3.180:8001
//*1A0  #continuo
//*1A1  #stop
//*1B1  #single
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
			line, err := reader.ReadString(CR)
			if err != nil {
				log.Fatal(err)
			}
			writer.WriteString(strings.TrimSpace(line))
			writer.WriteRune(CR)
			writer.WriteRune(NL)
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
		writer.WriteString(strings.TrimSpace(line))
		writer.WriteRune(CR)
		writer.WriteRune(NL)
		writer.Flush()
	}
}

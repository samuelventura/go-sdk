package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"

	"golang.org/x/net/proxy"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(os.Stdout)

	ctrlc := make(chan os.Signal, 1)
	signal.Notify(ctrlc, os.Interrupt)

	var proxyep = flag.String("p", "", "proxy endpoint (127.0.0.1:9999)")
	var address = flag.String("a", "", "address to connect to (192.168.0.254:80)")
	flag.Parse()

	if len(strings.TrimSpace(*proxyep)) == 0 ||
		len(strings.TrimSpace(*address)) == 0 {
		fmt.Printf("usage: %s -p 127.0.0.1:9999 -a 192.168.0.254:80\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	dialer, err := proxy.SOCKS5("tcp", *proxyep, nil, nil)
	if err != nil {
		log.Println(err)
		return
	}
	conn, err := dialer.Dial("tcp", *address)
	if err != nil {
		log.Println(err)
		return
	}

	done := make(chan interface{})
	go func() {
		io.Copy(os.Stdout, conn)
		done <- true
	}()
	go func() {
		io.Copy(conn, os.Stdin)
		done <- true
	}()
	select {
	case <-ctrlc:
	case <-done:
	}
}

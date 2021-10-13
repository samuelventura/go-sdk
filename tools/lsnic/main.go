package main

import (
	"fmt"
	"net"
)

//(cd tools/lsnic && go run .)
func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("localAddresses:", err.Error())
		return
	}
	for _, i := range ifaces {
		name := i.Name
		mac := i.HardwareAddr
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println("localAddresses:", err.Error())
			continue
		}
		for _, a := range addrs {
			fmt.Println(name, mac, a.Network(), a)
		}
	}
}

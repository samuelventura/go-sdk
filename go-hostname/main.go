package main

import (
	"fmt"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Hostname:", err.Error())
		return
	}
	fmt.Println(hostname)
}

package main

import (
	"log"
	"time"

	"github.com/goburrow/modbus"
)

//FIXME timeouts
//(cd go-comfile2; go run .)
func main() {
	handler := modbus.NewRTUClientHandler("/dev/ttyUSB0")
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 1
	handler.Timeout = 400 * time.Millisecond

	err := handler.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer handler.Close()

	client := modbus.NewClient(handler)
	for {
		_, err = client.WriteSingleCoil(4096, 0xff00)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
		_, err = client.WriteSingleCoil(4096, 0x0000)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
	}
}

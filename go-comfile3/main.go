package main

import (
	"log"
	"time"

	"github.com/samuelventura/go-modbus"
	"github.com/samuelventura/go-serial"
)

//FIXME timeouts
//(cd go-comfile3; go run .)
func main() {
	mode := &serial.Mode{}
	mode.BaudRate = 9600
	mode.DataBits = 8
	mode.Parity = serial.NoParity
	mode.StopBits = serial.OneStopBit
	trans, err := serial.NewSerialTransport("/dev/ttyUSB0", mode)
	if err != nil {
		log.Fatal(err)
	}
	defer trans.Close()
	trans.SetError(true)
	err = trans.Discard(200)
	if err != nil {
		log.Fatal(err)
	}
	modbus.EnableTrace(true)
	master := modbus.NewRtuMaster(trans, 4000)
	for {
		err = master.WriteDo(1, 4105, true)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
		err = master.WriteDo(1, 4105, false)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
	}
}

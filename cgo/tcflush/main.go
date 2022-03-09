package main

// #include <termios.h>
// static int get_TCIOFLUSH() { return TCIOFLUSH; }
import "C"
import (
	"fmt"
)

func main() {
	fmt.Println("TCIOFLUSH", C.get_TCIOFLUSH())
}

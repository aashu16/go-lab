package main

import (
	"fmt"
	"log"

	"github.com/jacobsa/go-serial/serial"
)

func jserial() {
	buf := make([]byte, 128)
	options := serial.OpenOptions{
		PortName: "COM4",
		BaudRate: 19200,
		DataBits: 8,
		StopBits: 1,
	}

	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("Open: %v", err)
	}

	defer port.Close()

	n, err := port.Write([]byte("*IDN?\n"))
	if err != nil {
		log.Fatalf("Write: %v", err)
	}

	r, err := port.Read(buf)
	if err != nil {
		log.Fatalf("Read: %v", err)
	}
	fmt.Println("IDN: ", r)

	fmt.Println("Wrote", n, "bytes")
}

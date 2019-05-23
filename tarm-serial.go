package main

import (
	"log"
	"time"

	"github.com/tarm/serial"
)

func tserial() {
	buf := make([]byte, 128)
	c := &serial.Config{Name: "COM4", Baud: 115200, ReadTimeout: 5 * time.Second}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n, err := s.Write([]byte("*IDN?\n"))
	if err != nil {
		log.Fatal(err)
	}

	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", buf[:n])
}

package serialPort

import (
	"log"
	"strings"
	"time"

	"github.com/tarm/serial"

	model "example/rhmtrz/txrx-mqtt/pkg/model"

	portList "go.bug.st/serial"
)

type SerialPort struct {
	serial *serial.Port
	// baudRate here?
}

func NewSerial() *SerialPort {
	// create SerialPort Struct
	// and then connect it
	return &SerialPort{serial: nil}
}

func (sp *SerialPort) Connect(baudRate int) {
	// retry to connect
	// if the port is connected, finish the function and return
	var s *serial.Port
	var err error
	for {
		ports, er := portList.GetPortsList()
		if er != nil {
			log.Println("no port list")
		}
		port := FilterPort(ports, func(p string) bool {
			return strings.Contains(p, "tty.usb") || strings.Contains(p, "ttyUSB0")
		})

		c := &serial.Config{Name: port, Baud: baudRate}
		s, err = serial.OpenPort(c)
		if err == nil {
			log.Println("Serial port connected!")
			break
		}
		// wait and log it
		log.Println("No serial port!")
		time.Sleep(2 * time.Second)

	}

	sp.serial = s
}

func (s *SerialPort) Read(buf []byte) int {
	// try to read with raw serial port
	// if it returns an error, retry connect

	n, err := (*s.serial).Read(buf)
	if err != nil {
		km := model.KeyManeger()
		keys := km.GetKeys()
		s.Close()
		s.Connect(keys.BaudRate)
	}
	return n

}

func (s *SerialPort) Write(buf []byte) int {
	// try to write with raw serial port
	// if it returns an error, retry connect
	n, err := (*s.serial).Write(buf)

	if err != nil {
		km := model.KeyManeger()
		keys := km.GetKeys()
		s.Close()
		s.Connect(keys.BaudRate)

	}
	return n
}

func (s *SerialPort) Close() {
	err := (*s.serial).Close()
	if err != nil {
		log.Println("Not closed", err)
	}
	// close the connection

}

func FilterPort(ports []string, isFindPort func(string) bool) string {
	var port string
	for _, p := range ports {
		if isFindPort(p) {
			port = p
		}
	}

	return port
}

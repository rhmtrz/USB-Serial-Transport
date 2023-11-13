package serialPort

import (
	"log"
	"strings"
	"time"

	"github.com/tarm/serial"
	serialPort "go.bug.st/serial"
)

type SerialPort struct {
	serial *serial.Port
}

func NewSerial() *SerialPort {
	return &SerialPort{serial: nil}
}

func (sp *SerialPort) Connect(baudRate int) {
	var s *serial.Port
	var err error

	for {
		ports, er := serialPort.GetPortsList()
		if er != nil {
			log.Println("No port list")
		}

		port := FilterPort(ports, func(p string) bool {
			return strings.Contains(p, "tty.usb") || strings.Contains(p, "ttyUSB0")
		})

		c := &serial.Config{Name: port, Baud: baudRate}
		s, err = serial.OpenPort(c)
		if err == nil {
			log.Println("Serial port connected...!")
			break
		}

		log.Println("No serial port..!")
		time.Sleep(2 * time.Second)
	}

	sp.serial = s
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

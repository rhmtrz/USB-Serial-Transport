package main

import (
	"flag"
	"log"

	model "./pkg/model"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	isLogging *bool
	version   *bool
)

func init() {
	isLogging = flag.Bool("logger", false, "bool")
	version = flag.Bool("version", false, "bool")
}

func main() {
	flag.Parse()
	if *version {
		log.Println("Version: 1.0.0")
		return
	}

	km := model.KeyManager()
	keys := km.GetKeys()
	println(keys.BaudRate)
	msgChan := make(chan mqtt.Message)

}

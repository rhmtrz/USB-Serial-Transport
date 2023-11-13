package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	model "example/rhmtrz/txrx-mqtt/pkg/model"
	"example/rhmtrz/txrx-mqtt/pkg/mqttConnection"
	"example/rhmtrz/txrx-mqtt/pkg/serialPort"

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
		log.Println("Version: 1.1")
		return
	}
	km := model.KeyManeger()
	keys := km.GetKeys()
	msgChan := make(chan mqtt.Message)
	serial := serialPort.NewSerial()
	serial.Connect(keys.BaudRate)

	clientManeger := mqttConnection.ClientManeger()
	cl := clientManeger.MqttConnect(keys.MqttURL, keys.DestinationId, msgChan)

	go clientManeger.PublishMqtt(keys.MyId, serial, *isLogging)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	for {
		select {
		case m := <-msgChan:
			if *isLogging {
				log.Println("sub: ", m.Payload(), m.Topic)
			}
			serial.Write(m.Payload())
		case <-signalCh:
			log.Println("!!!!!!!!!!!!")
			log.Printf("Interrupt detected.\n")
			cl.Disconnect(1000)
			return
		}
	}
}

package mqttConnection

import (
	"log"
	"time"

	"example/rhmtrz/txrx-mqtt/pkg/serialPort"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttClient struct {
	client *mqtt.Client
}

func ClientManeger() *MqttClient {
	return &MqttClient{client: nil}
}

func (mc *MqttClient) MqttConnect(url string, destinationId string, msgCh chan mqtt.Message) mqtt.Client {
	var client mqtt.Client

	opts := mc.CreateClientOptions(url, destinationId, msgCh)
	client = mqtt.NewClient(opts)
	token := client.Connect()
	token.Wait()
	if err := token.Error(); err != nil {
		log.Println("Please check your network connection")
		time.Sleep(3 * time.Second)
	}
	log.Println("network is avilable")

	mc.client = &client
	return *mc.client
}

func (c *MqttClient) GetClient() mqtt.Client {
	return *c.client
}

func (c *MqttClient) PublishMqtt(myId string, serial *serialPort.SerialPort, isLogging bool) {
	buf := make([]byte, 64)
	client := c.GetClient()
	for {
		n := serial.Read(buf)
		token := client.Publish(myId, 0, false, buf[:n])
		token.Wait()
		if isLogging {
			log.Println("Pub: ", buf[:n])
		}
	}
}

func (c *MqttClient) CreateClientOptions(url string, destinationId string, msgCh chan mqtt.Message) *mqtt.ClientOptions {
	var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		log.Println("Lost connection: ", err)
	}

	var onConnect mqtt.OnConnectHandler = func(client mqtt.Client) {
		log.Println("MQTT Connected !")
		subToken := client.Subscribe(destinationId, 0, nil)
		subToken.Wait()
		log.Println("Subscribed")

		if subToken.Error() != nil {
			log.Println("subscription err: ", subToken.Error())
		}
	}

	var messageHandler mqtt.MessageHandler = func(c mqtt.Client, msg mqtt.Message) {
		msgCh <- msg
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(url)
	opts.OnConnectionLost = connectLostHandler
	opts.DefaultPublishHandler = messageHandler
	opts.KeepAlive = 10
	opts.OnConnect = onConnect
	opts.SetAutoReconnect(true)
	return opts
}

package utilities

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type WarpKey struct {
	MqttURL       string `json:"mqttURL"`
	MyId          string `json:"myId"`
	DestinationId string `json:"destinationId"`
	BaudRate      int    `json:"baudRate"`
}

func KeyManeger() *WarpKey {
	return &WarpKey{}
}

func (k *WarpKey) LoadJson() {

	jsonFile, err := os.Open("./warpKey.json")
	defer jsonFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &k)
}

func (k *WarpKey) GetKeys() WarpKey {
	k.LoadJson()
	return *k
}

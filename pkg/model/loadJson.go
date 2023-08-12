package utilities

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type TxrxKey struct {
	MqttURL  string `json:"mqttURL"`
	TxId     string `json:"txId"`
	RxId     string `json:"rxId"`
	BaudRate int    `json:"baudRate"`
}

func KeyManager() *TxrxKey {
	return &TxrxKey{}
}

func (k *TxrxKey) LoadJson() {

	jsonFile, err := os.Open("./txrxKey.json")
	defer jsonFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &k)
}

func (k *TxrxKey) GetKeys() TxrxKey {
	k.LoadJson()
	return *k
}

package service

import (
	"encoding/json"
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type KasaConfiguration struct {
	DevicesPath string
	MQTTURL     string
}

type KasaDeviceGroup struct {
	PreferredName string   `json:"preferredName"`
	Names         []string `json:"names"`
	Devices       []string `json:"devices"`
}

type Kasa struct {
	client       mqtt.Client
	deviceGroups []KasaDeviceGroup
}

func NewKasa(configuration KasaConfiguration) (*Kasa, error) {
	devicesString, err := os.ReadFile(configuration.DevicesPath)
	if err != nil {
		return nil, err
	}

	var deviceGroups []KasaDeviceGroup
	err = json.Unmarshal(devicesString, &deviceGroups)
	if err != nil {
		return nil, err
	}

	return &Kasa{
		client:       mqtt.NewClient(mqtt.NewClientOptions().AddBroker(configuration.MQTTURL)),
		deviceGroups: deviceGroups,
	}, nil
}

func (k *Kasa) DeviceGroups() []KasaDeviceGroup {
	return k.deviceGroups
}

func (k *Kasa) DeviceNamesAndMap() ([]string, map[string]KasaDeviceGroup) {
	strings := make([]string, 0)
	deviceMap := make(map[string]KasaDeviceGroup)
	for _, device := range k.DeviceGroups() {
		for _, name := range device.Names {
			strings = append(strings, name)
			deviceMap[name] = device
		}
	}
	return strings, deviceMap
}

func (k *Kasa) SetStatus(id string, on bool) error {
	if !k.client.IsConnected() {
		if token := k.client.Connect(); token.Wait() && token.Error() != nil {
			return token.Error()
		}
	}

	var message string
	if on {
		message = "on"
	} else {
		message = "off"
	}
	topic := fmt.Sprintf("/%s/switch", id)
	token := k.client.Publish(topic, 0, false, message)
	token.Wait()
	return token.Error()
}

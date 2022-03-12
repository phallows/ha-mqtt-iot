package InternalDevice

import ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/external"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice BinarySensor) Translate() ExternalDevice.BinarySensor {
	eDevice := ExternalDevice.BinarySensor{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	return eDevice
}

type BinarySensor struct {
	State []string `json:"state"`
	MQTT  struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}

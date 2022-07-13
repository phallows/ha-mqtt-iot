package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
func (iDevice Button) Translate() externaldevice.Button {
	eDevice := externaldevice.Button{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.AvailabilityMode = iDevice.AvailabilityMode
	eDevice.AvailabilityTemplate = iDevice.AvailabilityTemplate
	eDevice.AvailabilityFunc = common.ConstructStateFunc(iDevice.Availability)
	eDevice.CommandTemplate = iDevice.CommandTemplate
	eDevice.CommandFunc = common.ConstructCommandFunc(iDevice.Command)
	eDevice.DeviceClass = iDevice.DeviceClass
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.Icon = iDevice.Icon
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.PayloadPress = iDevice.PayloadPress
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.UniqueId = iDevice.UniqueId
	if len(iDevice.Availability) == 0 {
		eDevice.AvailabilityFunc = common.AvailabilityFunc
	}
	eDevice.Initialize()
	return eDevice
}

type Button struct {
	AvailabilityMode     string   `json:"availability_mode"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate string   `json:"availability_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability         []string `json:"availability"`
	CommandTemplate      string   `json:"command_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `command_topic`."
	Command              []string `json:"command"`
	DeviceClass          string   `json:"device_class"`          // "The [type/class](/integrations/button/#device-class) of the button to set the icon in the frontend."
	EnabledByDefault     bool     `json:"enabled_by_default"`    // "Flag which defines if the entity should be enabled when first added."
	Encoding             string   `json:"encoding"`              // "The encoding of the published messages."
	EntityCategory       string   `json:"entity_category"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                 string   `json:"icon"`                  // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Name                 string   `json:"name"`                  // "The name to use when displaying this button."
	ObjectId             string   `json:"object_id"`             // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable     string   `json:"payload_available"`     // "The payload that represents the available state."
	PayloadNotAvailable  string   `json:"payload_not_available"` // "The payload that represents the unavailable state."
	PayloadPress         string   `json:"payload_press"`         // "The payload To send to trigger the button."
	Qos                  int      `json:"qos"`                   // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain               bool     `json:"retain"`                // "If the published message should have the retain flag on or not."
	UniqueId             string   `json:"unique_id"`             // "An ID that uniquely identifies this button entity. If two buttons have the same unique ID, Home Assistant will raise an exception."
	MQTT                 struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}

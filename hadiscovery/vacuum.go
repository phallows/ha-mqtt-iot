package hadiscovery

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (d Vacuum) GetRawId() string {
	return "vacuum"
}
func (d Vacuum) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d Vacuum) GetUniqueId() string {
	return d.UniqueId
}
func (d Vacuum) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Vacuum struct {
	AvailabilityMode     string                          `json:"availability_mode"`
	AvailabilityTemplate string                          `json:"availability_template"`
	AvailabilityTopic    string                          `json:"availability_topic"`
	AvailabilityFunc     func() string                   `json:"-"`
	CommandTopic         string                          `json:"command_topic"`
	CommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`
	Device               struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	Encoding            string                          `json:"encoding"`
	FanSpeedList        []string                        `json:"fan_speed_list"`
	Name                string                          `json:"name"`
	ObjectId            string                          `json:"object_id"`
	PayloadAvailable    string                          `json:"payload_available"`
	PayloadCleanSpot    string                          `json:"payload_clean_spot"`
	PayloadLocate       string                          `json:"payload_locate"`
	PayloadNotAvailable string                          `json:"payload_not_available"`
	PayloadPause        string                          `json:"payload_pause"`
	PayloadReturnToBase string                          `json:"payload_return_to_base"`
	PayloadStart        string                          `json:"payload_start"`
	PayloadStop         string                          `json:"payload_stop"`
	Qos                 int                             `json:"qos"`
	Retain              bool                            `json:"retain"`
	Schema              string                          `json:"schema"`
	SendCommandTopic    string                          `json:"send_command_topic"`
	SendCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	SetFanSpeedTopic    string                          `json:"set_fan_speed_topic"`
	SetFanSpeedFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	StateTopic          string                          `json:"state_topic"`
	StateFunc           func() string                   `json:"-"`
	SupportedFeatures   []string                        `json:"supported_features"`
	UniqueId            string                          `json:"unique_id"`
	MQTT                MQTTFields                      `json:"-"`
}

func (d Vacuum) UpdateState() {
	if d.AvailabilityTopic != "" {
		state := d.AvailabilityFunc()
		if state != stateStore.Vacuum.Availability[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.AvailabilityTopic, qos, retain, state)
			stateStore.Vacuum.Availability[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.StateTopic != "" {
		state := d.StateFunc()
		if state != stateStore.Vacuum.State[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.StateTopic, qos, retain, state)
			stateStore.Vacuum.State[d.UniqueId] = state
			token.Wait()
		}
	}
}
func (d Vacuum) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	if d.CommandTopic != "" {
		t := c.Subscribe(d.CommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.SendCommandTopic != "" {
		t := c.Subscribe(d.SendCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.SetFanSpeedTopic != "" {
		t := c.Subscribe(d.SetFanSpeedTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	token := c.Publish(GetDiscoveryTopic(d), 0, true, message)
	token.Wait()
	time.Sleep(500 * time.Millisecond)
	d.AnnounceAvailable()
	d.UpdateState()
}
func (d Vacuum) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, qos, retain, "offline")
	token.Wait()
	if d.CommandTopic != "" {
		t := c.Unsubscribe(d.CommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.SendCommandTopic != "" {
		t := c.Unsubscribe(d.SendCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.SetFanSpeedTopic != "" {
		t := c.Unsubscribe(d.SetFanSpeedTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d Vacuum) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, qos, retain, "online")
	token.Wait()
}
func (d Vacuum) Initialize() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d Vacuum) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		topicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.SendCommandFunc != nil {
		d.SendCommandTopic = GetTopic(d, "send_command_topic")
		topicStore[d.SendCommandTopic] = &d.SendCommandFunc
	}
	if d.SetFanSpeedFunc != nil {
		d.SetFanSpeedTopic = GetTopic(d, "set_fan_speed_topic")
		topicStore[d.SetFanSpeedTopic] = &d.SetFanSpeedFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}

package device

import (
	"encoding/json"
	"fmt"
	"github.com/brutella/hap/accessory"
	"home-hap/pkg/rest"
)

type shellyRestConfig struct {
	Ip    string
	Relay int
	Name  string
	Model string
}

type shellyRelayResponse struct {
	Ison     bool `json:"ison"`
	HasTimer bool `json:"has_timer"`
	// Unix timestamp of timer start; 0 if timer inactive or time not synced
	TimerStarted  int `json:"timer_started"`
	TimerDuration int `json:"timer_duration"`
}

func newShellyRest(cfg shellyRestConfig) *accessory.Lightbulb {
	a := accessory.NewLightbulb(accessory.Info{
		Name:         cfg.Name,
		SerialNumber: "10100",
		Manufacturer: "Shelly",
		Model:        cfg.Model,
		Firmware:     "20230107",
	})

	a.Lightbulb.On.OnValueRemoteUpdate(func(v bool) {
		go func() {
			client := rest.NewRestClient(cfg.Ip)
			turn := "off"
			if v {
				turn = "on"
			}

			path := fmt.Sprintf("/relay/%d?turn=%s", cfg.Relay, turn)
			response, _ := client.Get(path, nil)

			// read shelly response
			data := shellyRelayResponse{}
			json.Unmarshal(response, &data)

			// set status based on shelly response
			a.Lightbulb.On.SetValue(data.Ison)
		}()
	})

	return a
}

func NewShellyIngress(ip string) *accessory.Lightbulb {
	return newShellyRest(shellyRestConfig{
		Ip:    ip,
		Relay: 0,
		Name:  "Ingresso",
		Model: "1",
	})
}

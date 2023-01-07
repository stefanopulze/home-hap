package device

import (
	"fmt"
	"github.com/brutella/hap/accessory"
	"home-hap/internal/config"
	"home-hap/pkg/logging"
	"home-hap/pkg/rest"
	"strconv"
	"time"
)

type homeRelayRequest struct {
	Turn  string `json:"turn"`
	Timer uint16 `json:"timer"`
}

func NewGate(opts config.HomeRelay) *accessory.Door {
	a := accessory.NewDoor(accessory.Info{
		Name:         "Cancello",
		SerialNumber: "100",
		Firmware:     "0.0.1",
	})

	a.Door.TargetPosition.OnValueRemoteUpdate(func(v int) {
		l := logging.GetLoggerWithField("accessory", "gate")
		l.Info(fmt.Sprintf("new state is: %s", doorValueToString(v)))

		go func() {
			// 0 = close, 100 = open
			client := rest.NewRestClient(opts.Server)
			data := homeRelayRequest{Turn: "on", Timer: opts.Gate.Timer}

			client.PostJson("/relay/1", data)
		}()

		time.Sleep(10 * time.Second)
		// Set close as default, the command is like a switch
		a.Door.CurrentPosition.SetValue(0)
		a.Door.TargetPosition.SetValue(0)
	})

	return a
}

func NewIngressGate(opts config.HomeRelay) *accessory.Switch {
	a := accessory.NewSwitch(accessory.Info{
		Name:         "Cancello ingresso",
		SerialNumber: "100",
		Firmware:     "0.0.1",
	})

	a.Switch.On.OnValueRemoteUpdate(func(v bool) {
		l := logging.GetLoggerWithField("accessory", "ingress-gate")
		l.Info(fmt.Sprintf("new state gate: %s", booleanToString(v)))

		go func() {
			data := homeRelayRequest{
				Turn:  "on",
				Timer: opts.IngressGate.Timer,
			}

			client := rest.NewRestClient(opts.Server)
			if _, err := client.PostJson("/relay/0", data); err != nil {
				l.Error("Cannot change state of gate on relay 0")
			}
		}()

		time.Sleep(900 * time.Millisecond)
		a.Switch.On.SetValue(false)
	})

	return a
}

func doorValueToString(v int) string {
	if v == 0 {
		return "close"
	} else if v == 100 {
		return "open"
	} else {
		return strconv.Itoa(v)
	}
}

func booleanToString(v bool) string {
	if v {
		return "open"
	} else {
		return "close"
	}
}

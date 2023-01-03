package device

import (
	"fmt"
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/characteristic"
	"github.com/stefanopulze/daitem"
	"home-hap/internal/config"
	"home-hap/pkg/logging"
)

var client *daitem.Client

func NewDaitem(opts config.DaitemConfig) *accessory.SecuritySystem {
	cfg, _ := daitem.DefaultOptions(opts.Email, opts.Password, opts.MasterCode)
	client = daitem.NewClient(cfg)
	logger := logging.GetLoggerWithField("accessory", "daitem")

	a := accessory.NewSecuritySystem(accessory.Info{
		Name:         "Allarme",
		SerialNumber: "10010",
		Manufacturer: "Daitem",
		Firmware:     "0.0.5",
	})

	// Disable by default
	a.SecuritySystem.SecuritySystemTargetState.SetValue(characteristic.SecuritySystemTargetStateDisarm)
	a.SecuritySystem.SecuritySystemCurrentState.SetValue(characteristic.SecuritySystemTargetStateDisarm)

	a.SecuritySystem.SecuritySystemTargetState.OnValueRemoteUpdate(func(v int) {
		turnOn := v != characteristic.SecuritySystemTargetStateDisarm
		logger.Debug(fmt.Sprintf("Request new state: %d = %v", v, turnOn))

		go func() {
			if err := client.TurnAlarm(turnOn); err != nil {
				logger.Error(fmt.Sprintf("Cannot change remote state: %s", err))
				v = a.SecuritySystem.SecuritySystemCurrentState.Value()
				a.SecuritySystem.SecuritySystemTargetState.SetValue(v)
			}

			logger.Info(fmt.Sprintf("Change state in: %v", turnOn))

			a.SecuritySystem.SecuritySystemCurrentState.SetValue(v)
		}()
	})

	return a
}

package device

import "github.com/brutella/hap/accessory"

func NewBridge(name string) *accessory.Bridge {
	a := accessory.NewBridge(accessory.Info{
		Name:         name,
		SerialNumber: "101001",
		Manufacturer: "Stefano",
		Model:        "bridge",
		Firmware:     "0.0.2-alpha",
	})
	a.Id = 1

	return a
}

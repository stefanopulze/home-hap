package discovery

import (
	"context"
	"github.com/brutella/hap/accessory"
	"home-hap/internal/config"
	"home-hap/pkg/device"
)

type Discovery struct {
	ctx         context.Context
	cfg         *config.Config
	accessories []*accessory.A
}

func New(ctx context.Context, cfg *config.Config) *Discovery {
	return &Discovery{
		ctx: ctx,
		cfg: cfg,
	}
}

func (d *Discovery) LocalDevices() {
	ingressGate := device.NewIngressGate(d.cfg.HomeRelay)
	ingressGate.Id = 2
	d.accessories = append(d.accessories, ingressGate.A)

	gate := device.NewGate(d.cfg.HomeRelay)
	gate.Id = 3
	d.accessories = append(d.accessories, gate.A)

	daitem := device.NewDaitem(d.cfg.Daitem)
	daitem.Id = 4
	d.accessories = append(d.accessories, daitem.A)

	ingressLight := device.NewShellyIngress("http://192.168.20.51")
	ingressLight.Id = 5
	d.accessories = append(d.accessories, ingressLight.A)
}

func (d *Discovery) GetDevices() []*accessory.A {
	return d.accessories
}

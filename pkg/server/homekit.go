package server

import (
	"context"
	"fmt"
	"github.com/brutella/hap"
	"home-hap/internal/config"
	"home-hap/pkg/device"
	"home-hap/pkg/discovery"
	"home-hap/pkg/logging"
)

type Homekit struct {
	ctx    context.Context
	server *hap.Server
	logger *logging.Logger
}

func NewHomeKit(ctx context.Context, opts config.HomeKitOpts, disco *discovery.Discovery) *Homekit {
	logger := logging.GetLoggerWithField("server", "homekit")

	bridge := device.NewBridge(opts.Name)
	accs := disco.GetDevices()

	// Store the data in the "./db" directory.
	fs := hap.NewFsStore(opts.Storage)

	server, err := hap.NewServer(fs, bridge.A, accs...)
	if err != nil {
		// stop if an error happens
		logger.Error(err.Error())
		return nil
	}
	server.Pin = opts.Pin
	if len(opts.Ifaces) > 0 {
		logger.Debug(fmt.Sprintf("hap dns use custom ifaces: %+q", opts.Ifaces))
		server.Ifaces = opts.Ifaces
	}

	return &Homekit{
		ctx:    ctx,
		logger: logger,
		server: server,
	}
}

func (h *Homekit) Start() {
	h.logger.Info("Starting homekit server")
	go h.server.ListenAndServe(h.ctx)
}

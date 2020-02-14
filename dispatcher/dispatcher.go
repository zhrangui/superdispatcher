package dispatcher

import (
	"superdispatcher/config"
	"superdispatcher/logger"
	"superdispatcher/net"
)

// Dispatcher exposes network communication
type Dispatcher struct {
	config  *config.Config
	logger  *logger.Logger
	network *net.Network
}

// NewDispatcher creates service
func NewDispatcher(config *config.Config, logger *logger.Logger) (*Dispatcher, error) {
	dispatcher := &Dispatcher{
		config: config,
		logger: logger,
	}
	var err error
	dispatcher.network, err = net.NewNetwork(config, logger)
	return dispatcher, err
}

// Dispatch starts service
func (dispatcher *Dispatcher) Dispatch() {
	go dispatcher.listen()
}

func (dispatcher *Dispatcher) listen() {
	ln, err := dispatcher.network.Listen()
	if err != nil {
		dispatcher.logger.Error(err, "Failed to listen")
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			dispatcher.logger.Error(err, "Failed to accept")
			return
		}
		defer conn.Close()
	}
}

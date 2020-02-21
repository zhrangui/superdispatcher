package dispatcher

import (
	"superdispatcher/config"
	"superdispatcher/logger"
	"superdispatcher/net"
)

// Dispatcher exposes Net communication
type Dispatcher struct {
	*logger.Logger
	config *config.Config
	Net    *net.Net
}

// NewDispatcher creates service
func NewDispatcher(config *config.Config, logger *logger.Logger) (*Dispatcher, error) {
	dispatcher := &Dispatcher{
		config: config,
		Logger: logger,
	}
	var err error
	dispatcher.Net, err = net.NewNet(config)
	dispatcher.Net.Logger = logger

	return dispatcher, err
}

// Dispatch starts service
func (dispatcher *Dispatcher) Dispatch() {
	go dispatcher.listen()
}

func (dispatcher *Dispatcher) listen() {
	ln, err := dispatcher.Net.Listen()
	if err != nil {
		dispatcher.Logger.Error(err, "Failed to listen")
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			dispatcher.Logger.Error(err, "Failed to accept")
			return
		}
		defer conn.Close()
	}
}

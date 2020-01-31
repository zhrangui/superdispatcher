package dispatcher

import (
	"superdispatcher/config"
	"superdispatcher/logger"
	"superdispatcher/net"
)

// Dispatcher exposes network communication
type Dispatcher struct {
	config  *config.Config
	network *net.Network
	logger  *logger.Logger
}

// NewDispatcher creates service
func NewDispatcher(config *config.Config) (*Dispatcher, error) {
	var err error
	dispatcher := new(Dispatcher)
	dispatcher.network, err = net.NewNetwork(config)
	return dispatcher, err
}

// Dispatch starts service
func (dispatcher *Dispatcher) Dispatch() {
	go dispatcher.listen()
}

func (dispatcher *Dispatcher) listen() {
	ln, err := dispatcher.network.Listen()
	if err != nil {
		dispatcher.logger.Fatal(err.Error())
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			dispatcher.logger.Fatal(err.Error())
			return
		}
		defer conn.Close()
		if err != nil {
			dispatcher.logger.Fatal(err.Error())
		}
	}
}

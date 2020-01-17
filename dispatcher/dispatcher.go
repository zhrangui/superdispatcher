package dispatcher

import (
	"superdispatcher/config"
	"superdispatcher/net"
)

type Dispatcher struct {
	config  *config.Config
	network *net.Network
}

func New(config *config.Config) (*Dispatcher, error) {
	var err error
	dispatcher := new(Dispatcher)
	dispatcher.network, err = net.New(config)
	return dispatcher, err
}

func (dispatcher *Dispatcher) Dispatch() {

}

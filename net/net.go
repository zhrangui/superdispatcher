package net

import (
	"fmt"
	"net"

	"superdispatcher/config"
	"superdispatcher/logger"

	"github.com/pkg/errors"
)

// Net functionalities
type Net struct {
	config *config.Config
	*logger.Logger
}

// NewNet creates Net communication channel
func NewNet(config *config.Config) (*Net, error) {
	var net = &Net{
		config: config,
	}
	return net, nil
}

// Dial connects to a server
func (Net *Net) Dial() (net.Conn, error) {
	conn, err := net.Dial(Net.config.Constants.My.Protocal, Net.config.Constants.My.HostAddress)
	if err != nil {
		e := fmt.Sprintf("Fail to connect: %s", Net.config.Constants.My.HostAddress)
		Net.Logger.Fatal(e)
		err = errors.Wrap(err, e)
	}
	return conn, err
}

// Listen creates server
func (Net *Net) Listen() (net.Listener, error) {
	conn, err := net.Listen(Net.config.Constants.My.Protocal, Net.config.Constants.My.ServerIP)
	if err != nil {
		e := fmt.Sprintf("Fail to connect: %s", Net.config.Constants.My.ServerIP)
		Net.Logger.Fatal(e)
		err = errors.Wrap(err, e)
	}
	return conn, err
}

package net

import (
	"fmt"
	"net"

	"superdispatcher/config"

	"github.com/pkg/errors"
)

// Network functionalities
type Network struct {
	config *config.Config
}

// NewNetwork creates network communication channel
func NewNetwork(config *config.Config) (*Network, error) {
	var net = new(Network)
	net.config = config
	return net, nil
}

// Dial connects to a server
func (network *Network) Dial() (net.Conn, error) {
	conn, err := net.Dial(network.config.Constants.My.Protocal, network.config.Constants.My.HostAddress)
	if err != nil {
		e := fmt.Sprintf("Fail to connect: %s", network.config.Constants.My.HostAddress)
		network.config.Logger.Fatal(e)
		err = errors.Wrap(err, e)
	}
	return conn, err
}

// Listen creates server
func (network *Network) Listen() (net.Listener, error) {
	conn, err := net.Listen(network.config.Constants.My.Protocal, network.config.Constants.My.ServerIP)
	if err != nil {
		e := fmt.Sprintf("Fail to connect: %s", network.config.Constants.My.ServerIP)
		network.config.Logger.Fatal(e)
		err = errors.Wrap(err, e)
	}
	return conn, err
}

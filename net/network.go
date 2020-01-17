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

// New creates network communication channel
func New(config *config.Config) (*Network, error) {
	var net = new(Network)
	net.config = config
	return net, nil
}

// Listen creates connection to specified IP
func (network *Network) Listen() (net.Listener, error) {
	conn, err := net.Listen(network.config.Constants.My.Protocal, network.config.Constants.My.IpAddress)
	if err != nil {
		e := fmt.Sprintf("Fail to connect: %s", network.config.Constants.My.IpAddress)
		network.config.Logger.Logger.Fatal(e)
		err = errors.Wrap(err, e)
	}
	return conn, err
}

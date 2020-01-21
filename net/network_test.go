package net

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"superdispatcher/config"
)

func TestListenDial(t *testing.T) {
	//var err error
	config, err := config.New("config", "../resources")
	config.Constants.My.ServerIP = ":5000"
	config.Constants.My.HostAddress = ":5000"
	network, err := NewNetwork(config)
	go network.Listen()
	conn, err := network.Dial()
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	assert.NotNil(t, status, err)
}

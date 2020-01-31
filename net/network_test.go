package net

import (
	"fmt"
	"io/ioutil"

	"github.com/stretchr/testify/assert"
	"testing"

	"superdispatcher/config"
)

func TestListenDial(t *testing.T) {
	//var err error
	config, err := config.New("config", "../resources")
	network, err := NewNetwork(config)
	message := "dial test!\n"

	go func() {
		conn, err := network.Dial()
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		if _, err := fmt.Fprintf(conn, message); err != nil {
			t.Fatal(err)
		}
	}()

	ln, err := network.Listen()
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		defer conn.Close()

		buf, err := ioutil.ReadAll(conn)
		if err != nil {
			t.Fatal(err)
		}

		if msg := string(buf[:]); msg != message {
			assert.Equal(t, message, msg)
		}
		return
	}
}

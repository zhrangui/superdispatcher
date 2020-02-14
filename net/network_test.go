package net

import (
	"superdispatcher/config"
	"superdispatcher/logger"

	"fmt"
	"io/ioutil"
	"log"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListenDial(t *testing.T) {
	//var err error
	config, err := config.New("config", "../resources")
	if err != nil {
		log.Fatal(err)
	}
	logger, err := logger.NewLog(config)
	network, err := NewNetwork(config, logger)

	const (
		message = "dial test!"
	)

	go func() {
		conn, err := network.Dial()
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		if _, err := fmt.Fprintf(conn, message); err != nil {
			panic(err)
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
			t.Fatal(err)
			return
		}
		defer conn.Close()

		buf, err := ioutil.ReadAll(conn)
		if err != nil {
			t.Fatal(err)
		}

		msg := string(buf[:])
		assert.Equal(t, message, msg)
		return
	}
}

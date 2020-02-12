package main

import (
	"fmt"
	"log"

	"superdispatcher/config"
	"superdispatcher/dispatcher"
)

func main() {
	forever := make(chan bool)

	config, err := config.New("config", "resources")
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to init config: %+v", err))
	}

	dsptcher, err := dispatcher.NewDispatcher(config)

	if err != nil {
		log.Fatal(fmt.Sprintf("can't start service: %+v", err))
	}
	dsptcher.Dispatch()

	<-forever
}

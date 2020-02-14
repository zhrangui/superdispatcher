package main

import (
	"fmt"
	"log"

	"superdispatcher/config"
	"superdispatcher/dispatcher"
	"superdispatcher/logger"
)

func main() {
	forever := make(chan bool)

	config, err := config.New("config", "resources")
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to init config: %+v", err))
	}
	logger, err := logger.NewLog(config)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to init log: %+v", err))
	}
	dsptcher, err := dispatcher.NewDispatcher(config, logger)

	if err != nil {
		log.Fatal(fmt.Sprintf("can't start service: %+v", err))
	}
	dsptcher.Dispatch()

	<-forever
}

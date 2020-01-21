package main

import (
	"fmt"
	"log"
	"sync"

	"superdispatcher/config"
	"superdispatcher/dispatcher"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	config, err := config.New("config", "resources")
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to init config: %+v", err))
	}
	dsptcher, err := dispatcher.New(config)
	dsptcher.Dispatch()
	wg.Wait()
}

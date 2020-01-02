package main

import (
	"log"
	"sync"

	"superdispatcher/config"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	_, err := config.New("config", "resources")
	if err != nil {
		log.Fatal("failed to init config: %+v", err)
	}
	wg.Wait()
}

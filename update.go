package main

import (
	"net/http"

	"github.com/inconshreveable/go-update"

	"superdispatcher/config"
	"superdispatcher/logger"
)

func doUpdate(url string, config config.Config, logger logger.Logger) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		if rerr := update.RollbackError(err); rerr != nil {
			logger.Error(rerr, "Failed to rollback from bad update")
		}
	}
	return err
}

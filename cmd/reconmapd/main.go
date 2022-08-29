package main

import (
	log "github.com/sirupsen/logrus"
	"reconmap/agent/internal"
)

func main() {
	log.SetLevel(log.DebugLevel)

	app := internal.NewApp()
	if err := app.Run(); err != nil {
		log.Error(*err)
	}
}

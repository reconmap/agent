package main

import (
	"reconmap/agent/internal"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	app := internal.NewApp()
	app.Run()
}

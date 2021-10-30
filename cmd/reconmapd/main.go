package main

import (
	"reconmap/agent/internal"

)

func main() {
	log.SetLevel(log.DebugLevel)

	app := internal.NewApp()
	app.Run()
}

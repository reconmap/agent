package main

import (
	"reconmap/agent/internal"
)

func main() {
	app := &internal.App{}
	app.Initialise()
	app.Run()
}

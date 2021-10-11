package internal

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func handleNotifications(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling notification request")

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
	}

	registerClient(conn)
}

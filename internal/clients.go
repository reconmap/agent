package internal

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var clients []*websocket.Conn

func broadcast(message string) {
	log.Debug("broadcasting message")

	for _, client := range clients {
		log.Debug("-> " + client.RemoteAddr().String())
		client.WriteMessage(websocket.TextMessage, []byte(message))
	}
}

func registerClient(client *websocket.Conn) {
	log.Debug("registering client connection")
	clients = append(clients, client)
}

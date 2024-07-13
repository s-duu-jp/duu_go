// chat.go
package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[string]map[*websocket.Conn]bool)
var broadcast = make(chan Message)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Channel  string `json:"channel"`
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	channelID := r.URL.Query().Get("id")
	if channelID == "" {
		http.Error(w, "Missing channel ID", http.StatusBadRequest)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Upgrade: %v", err)
	}
	defer ws.Close()

	if _, ok := clients[channelID]; !ok {
		clients[channelID] = make(map[*websocket.Conn]bool)
	}
	clients[channelID][ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients[channelID], ws)
			break
		}
		msg.Channel = channelID
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		channelID := msg.Channel
		for client := range clients[channelID] {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients[channelID], client)
			}
		}
	}
}

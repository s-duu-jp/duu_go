// chat.go
package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket接続をアップグレードするための設定です。
// CheckOrigin関数は、すべてのオリジンからの接続を許可します。
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// clients: チャネルごとにWebSocket接続を管理するマップです。
// キーはチャネルIDで、値はそのチャネルに属するWebSocket接続のマップです。
var clients = make(map[string]map[*websocket.Conn]bool)

// broadcast: チャネルごとにメッセージをブロードキャストするためのチャネルです。
var broadcast = make(chan Message)

// Message: WebSocketで送受信されるメッセージの構造体です。
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Channel  string `json:"channel"`
}

// HandleConnectionsは、WebSocket接続を処理するハンドラーです。
func HandleConnections(w http.ResponseWriter, r *http.Request) {

	// クエリパラメータからチャネルIDを取得します。
	channelID := r.URL.Query().Get("id")
	// チャネルIDが空の場合は、400 Bad Requestを返します。
	if channelID == "" {
		http.Error(w, "Missing channel ID", http.StatusBadRequest)
		return
	}

	// WebSocket接続をアップグレードします。
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Upgrade: %v", err)
	}
	defer ws.Close()

	// チャネルごとにWebSocket接続を管理するためのマップを初期化します。
	if _, ok := clients[channelID]; !ok {
		clients[channelID] = make(map[*websocket.Conn]bool)
	}
	// WebSocket接続をマップに追加します。
	clients[channelID][ws] = true

	// メッセージを受信するための無限ループです。
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients[channelID], ws)
			break
		}
		// メッセージをブロードキャストします。
		msg.Channel = channelID
		// メッセージをブロードキャストするためのチャネルに送信します。
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		// メッセージを受信するための無限ループです。
		msg := <-broadcast
		// メッセージをチャネルごとにブロードキャストします。
		channelID := msg.Channel
		// チャネルに属するすべてのWebSocket接続にメッセージを送信します。
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

// chat.go
package websocket

import (
	"context"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
)

var ctx = context.Background()

// Redisクライアントの設定
var rdb = redis.NewClient(&redis.Options{
	Addr: "redis:6379",
})

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

	// 接続時に過去のメッセージを取得して送信する
	err = sendPastMessages(ws, channelID)
	if err != nil {
		log.Printf("Error sending past messages: %v", err)
	}

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
		// メッセージをRedisに保存する
		err = saveMessageToRedis(channelID, msg)
		if err != nil {
			log.Printf("Error saving message to Redis: %v", err)
		}
		// メッセージをブロードキャストするためのチャネルに送信します。
		broadcast <- msg
	}
}

// 過去のメッセージを送信する関数
func sendPastMessages(ws *websocket.Conn, channelID string) error {
	// Redisストリームから過去のメッセージを取得する
	messages, err := rdb.XRange(ctx, channelID, "-", "+").Result()
	if err != nil {
		return err
	}

	// 各メッセージをWebSocket接続に送信する
	for _, message := range messages {
		msg := Message{
			Username: message.Values["username"].(string),
			Message:  message.Values["message"].(string),
			Channel:  channelID,
		}
		err := ws.WriteJSON(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

// メッセージをRedisに保存する関数
func saveMessageToRedis(channelID string, msg Message) error {
	_, err := rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: channelID,
		Values: map[string]interface{}{
			"username": msg.Username,
			"message":  msg.Message,
		},
	}).Result()
	return err
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

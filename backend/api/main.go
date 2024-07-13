package main

import (
	sw "api/controllers/restapi"
	"api/handlers/restapi"
	"api/handlers/websocket"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// WebSocketのメッセージ処理をゴルーチンとして開始
	go websocket.HandleMessages()

	// ルーターにRESTful APIのルートを追加
	router := sw.NewRouter(sw.ApiHandleFunctions{
		AuthenticationAPI: sw.NewAuthenticationAPI(restapi.NewAuthenticationHandlers()),
	})

	// WebSocketのハンドラを追加
	router.GET("/chat", gin.WrapH(http.HandlerFunc(websocket.HandleConnections)))

	// サーバーの起動
	log.Fatal(router.Run(":3000"))
}

package server

import (
	"fmt"
	customWebsocket "kaustubh-upare/chatApplication/internals/Websocket"
	"log"
	"net/http"
)

func SetupRoutes() {
	fmt.Println("It is Working")
	pool := customWebsocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverWs(pool, w, r)
	})
}

func serverWs(pool *customWebsocket.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := customWebsocket.Upgrade(w, r)
	if err != nil {
		log.Println("error Occured in ServerWs")
		return
	}
	client := &customWebsocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	go client.Read()
}

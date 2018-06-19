package main

import (
	"golang.org/x/net/websocket"
	"github.com/helmedeiros/useless"
	"net/http"
)

func main() {
	http.Handle("/useless", websocket.Handler(func(ws *websocket.Conn) {
		ws.Write([]byte(useless.Gorillaz()))
	}))
	http.ListenAndServe(":3000", nil)
}

package Net

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type WsServer struct {
	url  string
	err  error
	Conn *websocket.Conn
}

func (ws WsServer) NewServer(w http.ResponseWriter, r *http.Request, responseHeader http.Header) {
	defer ws.Conn.Close()
	upd := websocket.Upgrader{}
	upd.EnableCompression = false
	ws.Conn, ws.err = upd.Upgrade(w, r, responseHeader)
}

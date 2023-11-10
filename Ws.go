package Net

import "github.com/gorilla/websocket"

type WebSocket struct {
	Conn *websocket.Conn
}

func (ws *WebSocket) NewWebsocketClient(url string) *WebSocket {
	if conn, _, err := websocket.DefaultDialer.Dial(url, nil); err != nil {
		ws.Conn = conn
	}
	return ws
}

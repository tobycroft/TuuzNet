package Net

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type WsServer struct {
	url          string
	err          error
	Conn         *websocket.Conn
	ReadChannel  chan []byte
	WriteChannel chan []byte
}

func (ws WsServer) prepare_channel() {
	ws.ReadChannel = make(chan []byte, 1)
	ws.WriteChannel = make(chan []byte, 1)
}

func (ws WsServer) NewServer(w http.ResponseWriter, r *http.Request, responseHeader http.Header) {
	defer ws.Conn.Close()
	upd := websocket.Upgrader{}
	upd.EnableCompression = false
	ws.Conn, ws.err = upd.Upgrade(w, r, responseHeader)
}

func (ws WsServer) recv_data() {
	for {
		_, message, err := ws.Conn.ReadMessage()
		if err != nil {
			ws.err = err
			log.Println("read:", err)
			return
		}
		ws.ReadChannel <- message
	}
}

func (ws WsServer) send_data() {
	for c := range ws.WriteChannel {
		err := ws.Conn.WriteMessage(websocket.TextMessage, c)
		if err != nil {
			ws.err = err
			log.Println("write:", err)
			return
		}
	}
}

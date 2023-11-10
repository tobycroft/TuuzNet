package Net

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

type WsData struct {
	Ip      string
	Conn    *websocket.Conn
	Message []byte
}

var WsServer_ReadChannel = make(chan WsData, 1)
var WsServer_WriteChannel = make(chan WsData, 1)

var client = new(sync.Map)

type WsServer struct {
	url  string
	err  error
	Conn *websocket.Conn
}

func (ws WsServer) NewServer(w http.ResponseWriter, r *http.Request, responseHeader http.Header) {
	defer ws.Conn.Close()
	upd := websocket.Upgrader{}
	upd.EnableCompression = false
	upd.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws.Conn, ws.err = upd.Upgrade(w, r, responseHeader)
	addr := ws.Conn.RemoteAddr()
	defer client.Delete(addr.String())
	client.Store(addr.String(), ws.Conn)
	go ws.send_data()
	for {
		_, message, err := ws.Conn.ReadMessage()
		if err != nil {
			ws.err = err
			log.Println("read:", err)
			return
		}
		WsServer_ReadChannel <- WsData{Ip: addr.String(), Conn: ws.Conn, Message: message}
	}
}

func (WsServer) send_data() {
	for c := range WsServer_WriteChannel {
		conn, ok := client.Load(c.Ip)
		if !ok {
			continue
		}
		err := conn.(*websocket.Conn).WriteMessage(websocket.TextMessage, c.Message)
		if err != nil {
			log.Println("senderror:", err)
			return
		}
	}
}

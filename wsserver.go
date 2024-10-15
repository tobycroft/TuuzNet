package Net

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

type WsData struct {
	Conn    *websocket.Conn
	Type    int
	Message []byte
	Status  bool
}

var WsServer_ReadChannel = make(chan WsData, 1)
var WsServer_WriteChannel = make(chan WsData, 1)

var WsConns sync.Map
var wsLock sync.Map

type WsServer struct {
	WsConfig *WsConfig
	url      string
	err      error
	Conn     *websocket.Conn
}

type WsConfig struct {
	PingReplyDelayInMs uint
	PongReplyDelayInMs uint
	Compress           bool
}

func (ws *WsServer) NewServer(w http.ResponseWriter, r *http.Request, responseHeader http.Header) {
	if ws.WsConfig == nil {
		ws.WsConfig = &WsConfig{
			PingReplyDelayInMs: 10,
		}
	}
	upd := websocket.Upgrader{
		EnableCompression: ws.WsConfig.Compress,
	}
	upd.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws.Conn, ws.err = upd.Upgrade(w, r, responseHeader)
	if ws.err != nil {
		return
	}
	go ws.send_data()
	defer ws.Conn.Close()
	WsConns.Store(ws.Conn.RemoteAddr().String(), ws.Conn)
	wsLock.Store(ws.Conn.RemoteAddr().String(), &sync.Mutex{})
	for {
		Type, message, err := ws.Conn.ReadMessage()
		switch Type {

		case websocket.TextMessage:
			WsServer_ReadChannel <- WsData{Conn: ws.Conn, Message: message, Type: Type, Status: true}
			break

		case websocket.BinaryMessage:
			WsServer_ReadChannel <- WsData{Conn: ws.Conn, Message: message, Type: Type, Status: true}
			break

		case websocket.PingMessage:
			go func() {
				time.Sleep(time.Duration(ws.WsConfig.PongReplyDelayInMs) * time.Millisecond)
				ws.Conn.WriteMessage(websocket.PongMessage, []byte("pong"))
			}()
			break

		case websocket.PongMessage:
			go func() {
				time.Sleep(time.Duration(ws.WsConfig.PingReplyDelayInMs) * time.Millisecond)
				ws.Conn.WriteMessage(websocket.PingMessage, []byte("ping"))
			}()
			break

		case websocket.CloseMessage, -1:
			WsConns.Delete(ws.Conn.RemoteAddr().String())
			wsLock.Delete(ws.Conn.RemoteAddr().String())
			go func() {
				select {
				case <-time.After(1 * time.Second):
					break
				case WsServer_WriteChannel <- WsData{Conn: ws.Conn, Message: message, Type: Type}:
					break
				}
			}()

			go func() {
				select {
				case <-time.After(1 * time.Second):
					break
				case WsServer_ReadChannel <- WsData{Conn: ws.Conn, Message: message, Type: Type}:
					break
				}
			}()
			return

		default:
			if err != nil {
				WsConns.Delete(ws.Conn.RemoteAddr().String())
				wsLock.Delete(ws.Conn.RemoteAddr().String())
				log.Println("server-read-error:", err)
				return
			}
			break
		}

	}
}

func (ws *WsServer) send_data() {
	for c := range WsServer_WriteChannel {
		switch c.Type {
		case websocket.TextMessage, websocket.BinaryMessage:
			lock, ok := wsLock.Load(c.Conn.RemoteAddr().String())
			if ok {
				lock.(*sync.Mutex).Lock()
				err := c.Conn.WriteMessage(c.Type, c.Message)
				lock.(*sync.Mutex).Unlock()
				if err != nil {
					WsConns.Delete(c.Conn.RemoteAddr().String())
					wsLock.Delete(c.Conn.RemoteAddr().String())
					log.Println("server-send-error:", err)
					return
				}
			}
			break

		case websocket.PingMessage:
			lock, ok := wsLock.Load(c.Conn.RemoteAddr().String())
			if ok {
				lock.(*sync.Mutex).Lock()
				err := c.Conn.WriteMessage(websocket.PingMessage, []byte("ping"))
				lock.(*sync.Mutex).Unlock()
				if err != nil {
					log.Println("server-ping-error:", err)
				}
			}
			break

		case websocket.PongMessage:
			lock, ok := wsLock.Load(c.Conn.RemoteAddr().String())
			if ok {
				lock.(*sync.Mutex).Lock()
				err := c.Conn.WriteMessage(websocket.PongMessage, []byte("pong"))
				lock.(*sync.Mutex).Unlock()
				if err != nil {
					log.Println("server-pong-error:", err)
				}
			}
			break

		case websocket.CloseMessage, -1:
			lock, ok := wsLock.Load(c.Conn.RemoteAddr().String())
			if ok {
				lock.(*sync.Mutex).Lock()
				err := c.Conn.WriteMessage(websocket.CloseMessage, []byte("close"))
				lock.(*sync.Mutex).Unlock()
				if err != nil {
					WsConns.Delete(c.Conn.RemoteAddr().String())
					wsLock.Delete(c.Conn.RemoteAddr().String())
					//log.Println("server-close-error:", err)
				}
			}
			return

		default:
			lock, ok := wsLock.Load(c.Conn.RemoteAddr().String())
			if ok {
				lock.(*sync.Mutex).Lock()
				err := c.Conn.WriteMessage(websocket.TextMessage, c.Message)
				lock.(*sync.Mutex).Unlock()
				if err != nil {
					WsConns.Delete(c.Conn.RemoteAddr().String())
					wsLock.Delete(c.Conn.RemoteAddr().String())
					log.Println("server-send-error:", err)
					return
				}
			}
			break
		}

	}
}

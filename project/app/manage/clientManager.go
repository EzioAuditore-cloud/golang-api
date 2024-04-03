package manage

import (
	"encoding/json"
	"time"

	"project/app/kafkaMQ"
	"project/middleWare/logger"
	general "project/model/General.go"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID         int
	UUID       string
	Addr       string
	Name       string
	Conn       *websocket.Conn
	Srv        *Server
	SendBytes  chan []byte
	RecvString chan string
	RecvBytes  chan general.ChatMessage
	State      int32 //0-online 1-offline
}

func NewClient(conn *websocket.Conn, srv *Server, user general.UserClient) *Client {
	client := &Client{
		ID:         user.ID,
		UUID:       user.UUID,
		Conn:       conn,
		Addr:       conn.RemoteAddr().String(),
		Name:       user.Name,
		SendBytes:  make(chan []byte),
		RecvString: make(chan string),
		RecvBytes:  make(chan general.ChatMessage),
		Srv:        srv,
		State:      0,
	}
	return client
}

func (c *Client) ListenSend() {
	for {
		select {
		case msg, ok := <-c.RecvBytes:
			if !ok {
				return
			}
			bytesMsg, err := json.Marshal(msg)
			if err != nil {
				logger.StructLog("Error", "ListenSend json.Marshal Error: %v", err)
			}
			c.Conn.WriteMessage(websocket.TextMessage, bytesMsg)
		}
	}
}

func (c *Client) Logout() {
	srv := c.Srv
	msg := general.StructreChatMsg("已下线", c.Name, c.ID, 0)
	srv.BroadCast(c, msg)
	srv.Clients.Delete(c.ID)
}

func (c *Client) DoMessage() {
	srv := c.Srv
	defer func() {
		c.Conn.Close()
	}()
	for {
		c.Conn.SetReadDeadline(time.Now().Add(time.Second * 100))
		n, message, err := c.Conn.ReadMessage()
		if n == 0 || err != nil {
			c.Logout()
			return
		}
		data := general.StructreChatMsg(string(message), c.Name, c.ID, 0)
		if data.SendToID > 0 {
			//TODO: 私聊
		} else {
			srv.BroadCast(c, data)
		}
		byteData, err := json.Marshal(data)
		if err != nil {
			logger.StructLog("Error", "Producer json.Marshal Error: %v", err)
		}
		kafkaMQ.Producer(c.ID, byteData)
	}
}

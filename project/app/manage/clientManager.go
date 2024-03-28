package manage

import (
	"time"

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
	RecvBytes  chan []byte
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
		RecvBytes:  make(chan []byte),
		Srv:        srv,
		State:      0,
	}
	return client
}

func (c *Client) ListenSend() {
	for {
		select {
		case bytesMsg, ok := <-c.RecvBytes:
			if !ok {
				return
			}
			c.Conn.WriteMessage(websocket.TextMessage, bytesMsg)
		}
	}
}

func (c *Client) Logout() {
	srv := c.Srv
	srv.BroadCast(c, []byte(c.Name+"已下线"))
	srv.Clients.Delete(c.ID)
}

func (c *Client) DoMessage() {
	srv := c.Srv
	defer func() {
		c.Conn.Close()
	}()
	for {
		c.Conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		n, message, err := c.Conn.ReadMessage()
		if n == 0 || err != nil {
			c.Logout()
			return
		}
		srv.BroadCast(c, message)
	}
}

package manage

import (
	"context"
	"fmt"
	"io"

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
		Srv:        srv,
		State:      0,
	}
	return client
}

func Login(conn *websocket.Conn, srv *Server, user general.UserClient) *Client {
	var client *Client
	if v, ok := srv.Clients.Load(user.ID); !ok {
		client = NewClient(conn, srv, user)
		srv.Clients.Store(user.ID, client)
	} else {
		client = v.(*Client)
		client.Srv = srv
	}
	go client.ListenSend()
	go client.SendMessage()
	return client
}

func (c *Client) ListenSend() {
	for {
		bytesMsg := <-c.RecvBytes
		c.Conn.WriteMessage(websocket.TextMessage, bytesMsg)
	}
}

func (c *Client) Logout() {
	srv := c.Srv
	srv.BroadCast(c, []byte("已下线"))
	c.State = 1
}

func (c *Client) SendMessage() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			c.Conn.Close()
			break
		}
		c.Srv.BroadcastChannel <- message
	}
}

func (c *Client) DoMessage(ctx context.Context, isLive chan bool) {
	srv := c.Srv
	for {
		select {
		case <-ctx.Done():
			c.Logout()
			srv.BroadCast(c, []byte("已下线"))
			srv.Clients.Delete(c.Addr)
			return
		default:
			n, message, err := c.Conn.ReadMessage()
			if err != nil && err != io.EOF {
				fmt.Println("conn.Read to buf err:", err)
				return
			}
			if n == 0 {
				c.Logout()
				return
			}
			srv.BroadCast(c, message)
			isLive <- true
		}
	}
}

package manage

import (
	"encoding/json"
	"log"
	"project/middleWare/logger"
	general "project/model/General.go"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Server struct {
	// Clients map[int]*Client
	Clients sync.Map
	//广播消息channel
	BroadcastChannel chan general.ChatMessage
	ID               int
}

var Srv = &Server{
	ID: 111,
	// Clients:          map[int]*Client{},
	BroadcastChannel: make(chan general.ChatMessage),
}

//监听广播channel
func (srv *Server) ListenMessage() {
	defer func() {
		p := recover()
		if p != nil {
			logger.StructLog("Error", "ListenMessage Err: %v", p)
			panic(p)
		}
	}()
	for {
		msg := <-srv.BroadcastChannel
		srv.Clients.Range(func(key, value any) bool {
			cli := value.(*Client)
			if cli.State == 0 {
				cli.RecvBytes <- msg
			}
			return true
		})
	}
}

func (srv *Server) BroadCast(c *Client, msg general.ChatMessage) {
	sendMsg := msg
	logger.StructLog("Info", "BroadCast:%v: %v", c.Addr, sendMsg)
	srv.BroadcastChannel <- msg
}

func (srv *Server) Handler(conn *websocket.Conn, user general.UserClient) {
	var client *Client
	if v, ok := srv.Clients.Load(user.ID); !ok {
		client = NewClient(conn, srv, user)
		srv.Clients.Store(user.ID, client)
	} else {
		client = v.(*Client)
		client.Srv = srv
	}
	log.Printf("用户量已达 %v ！", client.ID)
	msg := general.StructreChatMsg("已上线", client.Name, client.ID, 0)
	srv.BroadCast(client, msg)
	go client.ListenSend()
	go client.DoMessage()
	/******Stress  Test******/
	go func(*Client) {
		msg := general.ChatMessage{
			Content:  "Hello Hello Hello Hello Hello",
			SendToID: 0,
			FromID:   client.ID,
			FromName: client.Name,
		}
		sendMsg, err := json.Marshal(msg)
		if err != nil {
			logger.StructLog("Error", "Marshal:%v: %v", client.ID, err)
		}
		count := 0
		for {
			time.Sleep(1 * time.Second)
			err := client.Conn.WriteMessage(websocket.TextMessage, sendMsg)
			if err != nil {
				count++
				logger.StructLog("Error", "WriteMessage:%v: %v", client.ID, err)
			}
			if count > 10 {
				return
			}
		}
	}(client)
}

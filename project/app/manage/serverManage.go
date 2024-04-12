package manage

import (
	"project/middleWare/logger"
	general "project/model/General.go"
	"sync"

	"github.com/gorilla/websocket"
)

type Server struct {
	// Clients map[int]*Client
	Clients sync.Map
	//广播消息channel
	BroadcastChannel chan general.ChatMessage
	ID               int
	MsgCount         int
	SMutex           *sync.Mutex
	SRWMutex         *sync.RWMutex
}

var Srv = &Server{
	ID:       111,
	MsgCount: 0,
	// Clients:          map[int]*Client{},
	BroadcastChannel: make(chan general.ChatMessage, 10),
	SMutex:           &sync.Mutex{},
	SRWMutex:         &sync.RWMutex{},
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
		// log.Printf("监听阻塞")
		msg := <-srv.BroadcastChannel
		// log.Printf("已监听到消息: %v", msg)
		srv.Clients.Range(func(key, value any) bool {
			cli := value.(*Client)
			if cli.State == 0 {
				// log.Printf("消息发送阻塞")
				cli.RecvBytes <- msg
				// log.Printf("消息%v已发送给%v", msg, cli.Name)
			}
			return true
		})
	}
}

func (srv *Server) BroadCast(c *Client, msg general.ChatMessage) {
	defer func() {
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	// sendMsg := msg
	// logger.StructLog("Info", "BroadCast:%v: %v", c.Addr, sendMsg)
	// index := c.ID % len(srv.BroadcastChannel)
	// log.Printf("%v 消息阻塞", c.ID)
	srv.BroadcastChannel <- msg
	// log.Printf("%v 消息已进入管道", c.ID)
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
	// log.Printf("用户量已达 %v ！", client.ID)
	msg := general.StructreChatMsg("已上线", client.Name, client.ID, 0)
	srv.BroadCast(client, msg)
	go client.ListenSend()
	go client.DoMessage()
	/******Stress  Test******/
	// go func(*Client) {
	// 	defer func() {
	// 		if p := recover(); p != nil {
	// 			panic(p)
	// 		}
	// 	}()
	// 	for {
	// 		if client.State == 1 {
	// 			return
	// 		}
	// 		// client.Srv.SMutex.Lock()
	// 		log.Printf("SMutex已加锁")
	// 		// client.Srv.MsgCount++
	// 		msg := general.ChatMessage{
	// 			Content:  client.Name + ": Hello Hello Hello Hello Hello",
	// 			SendToID: 0,
	// 			FromID:   client.ID,
	// 			FromName: client.Name,
	// 		}
	// 		client.Srv.BroadCast(client, msg)
	// 		// srv.SMutex.Unlock()
	// 		log.Printf("SMutex锁已释放")
	// 		byteData, err := json.Marshal(msg)
	// 		if err != nil {
	// 			logger.StructLog("Error", "Producer json.Marshal Error: %v", err)
	// 		}
	// 		kafkaMQ.Producer(client.ID, byteData)
	// 		log.Printf("成功发送 %v", client.Srv.MsgCount)
	// 		time.Sleep(10 * time.Second)

	// 	}
	// }(client)
}

package manage

import (
	"fmt"
	"project/middleWare/logger"
	general "project/model/General.go"
	"sync"

	"github.com/gorilla/websocket"
)

type Server struct {
	// Clients map[int]*Client
	Clients sync.Map
	//广播消息channel
	BroadcastChannel chan []byte
	ID               int
}

var Srv = &Server{
	ID: 111,
	// Clients:          map[int]*Client{},
	BroadcastChannel: make(chan []byte),
}

//监听广播channel
func (srv *Server) ListenMessage() {
	// defer func() {
	// 	p := recover()
	// 	if p != nil {
	// 		logger.StructLog("Error", "ListenMessage Err: %v", p)
	// 		panic(p)
	// 	}
	// }()
	for {
		fmt.Println("ListenMessage:", srv.ID)
		msg := <-srv.BroadcastChannel
		srv.Clients.Range(func(key, value any) bool {
			cli := value.(*Client)
			if cli.State == 0 {
				cli.RecvBytes <- msg
			}
			return true
		})
		// for _, v := range srv.Clients {
		// 	if v.State == 0 {
		// 		v.RecvBytes <- msg
		// 	}
		// }
	}
}

func (srv *Server) LoginSrv(conn *websocket.Conn, user general.UserClient) error {
	// connAddr := ""
	// conn, ok := peer.FromContext(ctx)
	// if ok {
	// 	connAddr = conn.Addr.String()
	// }
	// connAddr := conn.RemoteAddr().String()
	// defer func() {
	// 	e := recover()
	// 	if e != nil {
	// 		fmt.Println("panic: ", e)
	// 	}
	// 	fmt.Printf("%v连接断开\n", connAddr)
	// }()
	// client := Login(conn, srv, user)
	// srv.BroadCast(client, []byte("已上线"))
	// fmt.Printf("%v连接成功\n", client.ID)
	// isLive := make(chan bool)
	// ctx, cancel := context.WithCancel(context.Background())
	// go func(context.Context) {
	// 	client.DoMessage(ctx, isLive)
	// }(ctx)
	// for {
	// 	select {
	// 	case <-isLive:
	// 	case <-time.After(time.Second * 100):
	// 		cancel()
	// 	}
	// }
	return nil
}

func (srv *Server) BroadCast(c *Client, msg []byte) {
	sendMsg := "[" + c.Addr + "]" + c.UUID + ":" + string(msg)
	logger.StructLog("Info", "BroadCast:%v: %v", c.Addr, sendMsg)
	srv.BroadcastChannel <- msg
}

func (srv *Server) Handler(conn *websocket.Conn, user general.UserClient) {
	client := Login(conn, srv, user)
	// defer func() {
	// 	e := recover()
	// 	if e != nil {
	// 		fmt.Println("panic: ", e)
	// 	}
	// 	conn.Close()
	// 	fmt.Printf("%v连接断开\n", client.ID)
	// }()

	fmt.Printf("%v连接成功\n", client.ID)

	// isLive := make(chan bool)
	// ctx, cancel := context.WithCancel(context.Background())
	// go func(context.Context) {
	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			client.Logout()
	// 			srv.BroadCast(client, []byte("已下线"))
	// 			srv.Clients.Delete(client.ID)
	// 			return
	// 		default:
	// 			n, message, err := client.Conn.ReadMessage()
	// 			if n == 0 {
	// 				srv.BroadCast(client, []byte("已下线"))
	// 				srv.Clients.Delete(client.ID)
	// 				return
	// 			}
	// 			if err != nil && err != io.EOF {
	// 				fmt.Println("conn.Read to buf err:", err)
	// 				return
	// 			}
	// 			srv.BroadCast(client, message)
	// 			isLive <- true
	// 		}
	// 	}
	// }(ctx)
	// for {
	// 	select {
	// 	case <-isLive:
	// 	case <-time.After(time.Second * 100):
	// 		cancel()
	// 		return
	// 	}
	// }
}

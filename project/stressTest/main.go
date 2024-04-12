package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type user struct {
	id   int
	name string
}

type tokenString struct {
	token string
}

func main() {
	count := 1
	// wg := sync.WaitGroup{}
	for {
		// msg := general.ChatMessage{
		// 	Content:  "hello",
		// 	SendToID: 0,
		// 	FromID:   count,
		// 	FromName: "ezio",
		// }
		// wg.Add(1)
		go func(int) {
			// defer wg.Done()
			fmt.Printf("%v 连接\n", count)
			t := LoginTest(count)
			connServer(t)
		}(count)
		time.Sleep(time.Millisecond * 500)
		count++
		// if count == 30 {
		// 	break
		// }
		// wg.Wait()
		// return
	}
	time.Sleep(time.Minute * 10)
}

func LoginTest(count int) string {
	data := url.Values{}
	data.Set("id", strconv.Itoa(count))
	data.Set("userName", "ezio"+strconv.Itoa(count))
	url := "http://127.0.0.1:8080/login"
	res, err := http.PostForm(url, data)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return ""
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("33333", err.Error())
		return ""
	}
	t := strings.Split(string(content), ":")[1]
	t = t[1 : len(t)-2]
	// fmt.Println(t)
	return t
}

func connServer(t string) {
	url := "http://127.0.0.1:8080/conn"
	req, _ := http.NewRequest("GET", url, nil)
	// dialer := &websocket.Dialer{}
	// conn, resp, err := dialer.Dial("ws://127.0.0.1:3434/add", nil)
	// conn
	req.Header.Set("Authorization", t)
	req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Upgrade", "websocket")
	req.Header.Set("Sec-WebSocket-Version", "13")
	req.Header.Set("Sec-WebSocket-Key", "x3JJHMbDL1EzLkh9GBhXDw==")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("Websocket Conn Error:%v", err)
		return
	}
	defer resp.Body.Close()
	time.Sleep(10 * time.Minute)
}

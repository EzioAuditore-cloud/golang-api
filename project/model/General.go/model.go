package general

import (
	"time"
)

type UserClient struct {
	ID   int
	Name string
	UUID string
}

type ChatMessage struct {
	Content  string `json:"content"`
	SendToID int    `json:"send_to_id"`
	FromID   int    `json:"from_id"`
	FromName string `json:"from_name"`
	SendTime int64  `json:"send_time"`
	State    int32  `json:"state"`
}

func StructreChatMsg(content, fromName string, fromID, sendTo int) ChatMessage {
	chatMsg := ChatMessage{
		Content:  content,
		SendToID: sendTo,
		FromID:   fromID,
		FromName: fromName,
		SendTime: time.Now().Unix(),
		State:    0,
	}
	return chatMsg
}

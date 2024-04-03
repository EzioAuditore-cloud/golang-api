package dbModel

import (
	orm "project/database"
	general "project/model/General.go"
	"time"
)

type ChatMessageRecord struct {
	ID       int64  `json:"id"`
	Content  string `json:"content"`
	SendToID int    `json:"send_to_id"`
	FromID   int    `json:"from_id"`
	SendTime int64  `json:"send_time"`
	State    int32  `json:"state"`
}

func StructreChatMsgRecord(chatMsg general.ChatMessage) ChatMessageRecord {
	chatRecord := ChatMessageRecord{
		Content:  chatMsg.Content,
		SendToID: chatMsg.SendToID,
		FromID:   chatMsg.FromID,
		SendTime: time.Now().Unix(),
		State:    chatMsg.State,
	}
	return chatRecord
}

func (chatMsg ChatMessageRecord) InsertChatMsg() (chatID int64, err error) {
	result := orm.Db.Create(&chatMsg)
	chatID = chatMsg.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (chatMsg ChatMessageRecord) ListChatMsg(chatID int64, err error) {
	result := orm.Db.Create(&chatMsg)
	chatID = chatMsg.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

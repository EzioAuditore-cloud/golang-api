package kafkaMQ

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"project/middleWare/logger"
	general "project/model/General.go"

	"github.com/IBM/sarama"
)

func Consumer() {
	config := sarama.NewConfig()
	client, err := sarama.NewConsumerGroup([]string{"127.0.0.1:9092"}, chatTopic+"_group", config)
	if err != nil {
		panic(err)
	}
	go func(sarama.ConsumerGroup) {
		for {
			err = client.Consume(context.Background(), []string{chatTopic}, &handler{})
		}
	}(client)

}

type handler struct{}

func (h *handler) Setup(session sarama.ConsumerGroupSession) error {
	fmt.Println("setup")
	return nil
}

func (h *handler) Cleanup(session sarama.ConsumerGroupSession) error {
	fmt.Println("cleanup")
	return nil
}

func (h *handler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		// fmt.Printf("Message topic:%q partition:%d offset:%d value:%v\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
		chatMsg := general.ChatMessage{}
		if err := json.Unmarshal(msg.Value, &chatMsg); err != nil {
			logger.StructLog("Error", "ConsumeClaim json.Unmarshal Error: %v", err)
			continue
		}
		log.Println(chatMsg)
		// 自动提交
		session.MarkMessage(msg, "")
	}
	return nil
}
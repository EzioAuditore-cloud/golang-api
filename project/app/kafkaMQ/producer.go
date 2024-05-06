package kafkaMQ

import (
	"fmt"
	"os"
	"project/middleWare/logger"
	"time"

	"github.com/IBM/sarama"
)

func ProducerConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Timeout = 1 * time.Second
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	return config
}

func Producer(clientID int, byteData []byte) {
	topic := chatTopic
	config := ProducerConfig()
	addrs := os.Getenv("KAFKA_BROKERS")
	producer, err := sarama.NewSyncProducer([]string{addrs}, config)
	if err != nil {
		logger.StructLog("Error", "NewSyncProducer Error: %v", err)
		return
	}
	defer producer.Close()
	producerMsg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(fmt.Sprintf("%d", clientID)),
		Value: sarama.ByteEncoder(byteData),
	}
	_, _, err1 := producer.SendMessage(producerMsg)
	if err1 != nil {
		logger.StructLog("Error", "SendMessage Error: %v", err)
		return
	}
}

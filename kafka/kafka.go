package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
)

var KafkaProducer = KafkaConnectProducer()

func KafkaConnectProducer() sarama.SyncProducer {
	brokersUrl := []string{"localhost:29092"}
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		panic(err.Error())
	}

	return conn
}

func PushMsgToQueue(topic string, message []byte) error {

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := KafkaProducer.SendMessage(msg)

	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	return nil
}

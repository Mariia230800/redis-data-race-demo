package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/Mariia230800/redis-data-race-demo/internal/log"
)

type Producer interface {
	SendMessage(topic string, key string, value []byte) error
	Close() error
}

type KafkaProducer struct {
	syncProducer sarama.SyncProducer
}

func NewKafkaProducer(brokers []string) (Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 3

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka producer: %w", err)
	}

	return &KafkaProducer{syncProducer: producer}, nil
}

func (p *KafkaProducer) SendMessage(topic string, key string, value []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(value),
	}

	partition, offset, err := p.syncProducer.SendMessage(msg)
	if err != nil {
		log.Errorf("failed to send message to kafka: %v", err)
		return fmt.Errorf("send message: %w", err)
	}

	log.Infof("Message sent to Kafka topic=%s partition=%d offset=%d",
		topic, partition, offset)

	return nil
}

func (p *KafkaProducer) Close() error {
	return p.syncProducer.Close()
}

package kafka_service

import "context"

type KafkaService struct {
	Ctx context.Context
}

func SendToKafka(topic string, data []byte) {

}

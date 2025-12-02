package modules

type Kafka struct {
	KafkaBroker string `env:"KAFKA_BROKER" envDefault:"kafka:9092"`
	Topic       string `env:"KAFKA_TOPIC" envDefault:"movies-topic"`
}

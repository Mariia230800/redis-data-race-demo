package modules

type Kafka struct {
	KafkaBroker string `env:"KAFKA_BROKER" envDefault:"localhost:9092"`
}

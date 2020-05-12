package configuration

import (
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	flagForKafkaHost = "kafka_host"
	flagForKafkaPort = "kafka_port"
)

type KafkaConfig struct {
	Host string
	Port int
}

func LoadKafkaConfig() KafkaConfig {
	c := KafkaConfig{
		Host: "localhost",
		Port: 9092,
	}

	flag.String(flagForKafkaHost, c.Host, "Kafka host")
	flag.Int(flagForKafkaPort, c.Port, "Kafka port")

	flag.Parse()

	viper.BindPFlag(flagForKafkaHost, flag.Lookup(flagForKafkaHost))
	viper.BindPFlag(flagForKafkaPort, flag.Lookup(flagForKafkaPort))

	viper.AutomaticEnv()

	c.Host = viper.GetString(flagForKafkaHost)
	c.Port = viper.GetInt(flagForKafkaPort)

	return c
}

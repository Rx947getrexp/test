package initialize

import (
	"github.com/Shopify/sarama"
	"go-speed/global"
)

func initKafkaClient() sarama.Client {
	if len(global.Config.Kafka.Addr) == 0 {
		global.Logger.Warn().Msg("kafka未配置")
		return nil
	}
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Fetch.Max = 100_000_000 // 100m
	config.Producer.Return.Successes = true
	config.Producer.MaxMessageBytes = 100_000_000 // 100m
	client, err := sarama.NewClient(global.Config.Kafka.Addr, config)
	if err != nil {
		global.Logger.Err(err).Msg("连接kafka")
		return nil
	}
	return client
}

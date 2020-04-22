package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

var KSProductClient sarama.SyncProducer

func ProductInit() {
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true
	var err error
	// 使用给定代理地址和配置创建一个同步生产者
	KSProductClient, err = sarama.NewSyncProducer([]string{viper.GetString("kafka.ConsumerAddr")}, config)
	if err != nil {
		logrus.Error("[ProductInit] err msg", err.Error())
	}

}

func ProductStart(topic, value string) {
	syncProducer(topic, value)
	//asyncProducer1(Address)
}

//同步消息模式
func syncProducer(topic, value string) {
	for i := 0; i < 10; i++ {
		value := fmt.Sprintf(value, i)
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(value),
		}
		part, offset, err := KSProductClient.SendMessage(msg)
		if err != nil {
			log.Printf("send message(%s) err=%s \n", value, err)
		} else {
			fmt.Fprintf(os.Stdout, value+"发送成功，partition=%d, offset=%d \n", part, offset)
		}
		time.Sleep(2 * time.Second)
	}
}

package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var KConsumerClient sarama.Consumer

func ConsumerInit() {
	//配置
	config := sarama.NewConfig()
	//接收失败通知
	config.Consumer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	config.Version = sarama.V2_4_0_0
	//新建一个消费者
	var err error
	KConsumerClient, err = sarama.NewConsumer([]string{viper.GetString("kafka.ConsumerAddr")}, config)
	if err != nil {
		logrus.Error("[ConsumerInit] err msg", err.Error())
	}

}

func ConsumerStart(topic string) string {
	//根据消费者获取指定的主题分区的消费者,Offset这里指定为获取最新的消息.
	partitionConsumer, err := KConsumerClient.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("[ConsumerStart] error get partition consumer", err)
	}
	defer partitionConsumer.Close()
	//循环等待接受消息.
	for {
		select {
		//接收消息通道和错误通道的内容.
		case msg := <-partitionConsumer.Messages():
			fmt.Println("msg offset: ", msg.Offset, " partition: ", msg.Partition, " timestrap: ", msg.Timestamp.Format("2006-Jan-02 15:04"), " value: ", string(msg.Value))
		case err := <-partitionConsumer.Errors():
			fmt.Println(err.Err)
		}
	}
}

package mq

import (
	"fmt"

	"time"

	"strings"

	"github.com/Shopify/sarama"
	conf "github.com/nikugame/public/config"
)

type MsgContent struct {
	Content string //内容
	MsgType string //消息类型,充值、登录等等
	Key     string
}

type Client struct {
	Addrs       string
	Producer    sarama.SyncProducer
	AsyncEnable bool
	config      *sarama.Config
	MsgContent  MsgContent
}

func (s *Client) init() {
	s.config = sarama.NewConfig()
	s.config.Producer.RequiredAcks = sarama.WaitForAll
	s.config.Producer.Partitioner = sarama.NewRandomPartitioner
	if len(s.Addrs) == 0 {
		//logs.Info("没有设置kafka的主机或端口，使用配置文件。")
		cf, _ := conf.NewConfig("ini", "conf/settings.conf")
		s.Addrs = cf.String("MqServer::addrs")
		//fmt.Println(s.Addrs)
		// s.connString = cf.String("MqServer::Host") + ":" + cf.String("MqServer::Port")
		//logs.Info("连接字符串为：%s", s.Addrs)
	}
}

//Send 新建一个对象
func (s *Client) Send() (int32, int64, error) {
	s.init()
	var err error
	s.Producer, err = sarama.NewSyncProducer(strings.Split(s.Addrs, ","), s.config)
	if err != nil {
		fmt.Println(err)
	}
	defer s.Producer.Close()
	return s.Producer.SendMessage(s.msg())
}

func (s *Client) msg() *sarama.ProducerMessage {
	var msg sarama.ProducerMessage
	if s.MsgContent.MsgType != "" {
		msg.Topic = s.MsgContent.MsgType
	} else {
		fmt.Println("Topic is empty!")
	}
	// msg.Partition = int32(-1)
	msg.Timestamp = time.Now().Local()
	if s.MsgContent.Key != "" {
		msg.Key = sarama.StringEncoder(s.MsgContent.Key)
	}

	if s.MsgContent.Content != "" {
		msg.Value = sarama.ByteEncoder(s.MsgContent.Content)

	}
	return &msg
}

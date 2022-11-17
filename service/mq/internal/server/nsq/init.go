package nsq

import (
	"errors"
	"essrv/service/mq/internal/logic"
	"essrv/service/mq/internal/svc"

	mc "essrv/service/mq/internal/config"

	"github.com/nsqio/go-nsq"
)

type MsgHandler struct {
	SrvName string
}

func (m *MsgHandler) HandleMessage(msg *nsq.Message) (err error) {
	err = logic.SyncDataToEsLogic(string(msg.Body))
	if err != nil {
		return err
	}
	// fmt.Printf("%s recv from %v, msg:%v\n", m.SrvName, msg.NSQDAddress, string(msg.Body))
	return nil
}

func InitConsumer(topic string, channel string, address string) (err error) {
	// config := nsq.NewConfig()
	// config.LookupdPollInterval = 15 * time.Second
	// c, err := nsq.NewConsumer(topic, channel, config)
	c, _ := svc.NewNsqClient(topic, channel)
	c.AddHandler(&MsgHandler{ // or AddConcurrentHandlers
		SrvName: mc.SrvName,
	})
	if err := c.ConnectToNSQD(address); err != nil { // or ConnectToNSQLookupd
		return errors.New("ConnectToNSQD fail")
	}
	return nil

}

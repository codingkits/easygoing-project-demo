package svc

import (
	mc "essrv/service/mq/internal/config"

	"github.com/nsqio/go-nsq"
)

type ServiceContext struct {
	NsqConsumer *nsq.Consumer
}

func NewServiceContext() *ServiceContext {
	nc, _ := NewNsqClient(mc.NsqTopic, mc.NsqChan)
	svc := &ServiceContext{
		NsqConsumer: nc,
	}
	return svc
}

package svc

import (
	"errors"
	"time"

	"github.com/nsqio/go-nsq"
)

func NewNsqClient(topic string, channel string) (*nsq.Consumer, error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	nc, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return nc, errors.New("create consumer failed")
	}
	return nc, err
}

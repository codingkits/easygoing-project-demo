package main

import (
	"essrv/service/mq/internal/config"
	"essrv/service/mq/internal/server/nsq"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// svc := svc.ServiceContext().NsqConsumber
	err := nsq.InitConsumer(config.NsqTopic, config.NsqChan, config.NsqD1Add)
	if err != nil {
		fmt.Printf("init consumer failed, err:%v \n", err)
		return
	}
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	<-c
}

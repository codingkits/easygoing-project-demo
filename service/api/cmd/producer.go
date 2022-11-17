package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	nsq "github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func initProducer(str string) (err error) {
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(str, config)
	if err != nil {
		fmt.Printf("create producer failed, err:%v\n", err)
		return err
	}
	return nil
}

func mockMsg() string {
	m := map[string]string{
		"idfa":     "idfa_" + RandomString(1),
		"idfa_md5": "md5_idfa_" + RandomString(1),
		"imei":     "imei_" + RandomString(1),
		"imei_md5": "imei_md5_" + RandomString(1),
		"oaid":     "oaid_" + RandomString(1),
		"oaid_md5": "oaid_md5_" + RandomString(1),
	}
	mj, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
	}
	return string(mj)
}

func main() {
	nsqAddress := "192.168.199.77:4150"
	err := initProducer(nsqAddress)
	if err != nil {
		fmt.Printf("init producer failed, err:%v\n", err)
		return
	}
	for {
		err = producer.Publish("es_topic", []byte(mockMsg()))
		if err != nil {
			fmt.Println("pub fail")
		}
		fmt.Println("\n pub msg:", mockMsg())
		time.Sleep(time.Millisecond * 3000)
	}
}

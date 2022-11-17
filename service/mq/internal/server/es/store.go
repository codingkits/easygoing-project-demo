package es

import (
	"context"
	"essrv/service/mq/internal/config"
	"essrv/service/mq/internal/types"

	"github.com/olivere/elastic/v7"
)

// 插入:结构体方式插入
func StoreByStruct(client *elastic.Client, d types.DeviceInfo) error {
	_, err := client.Index().Index(config.EsIndexName).BodyJson(d).Do(context.Background())
	if err != nil {
		return err
	}
	// fmt.Printf("indexed %s to index %s, type %s \n", put.Id, put.Index, put.Type)
	return nil
}

func StoreByJsonStr(client *elastic.Client, js string) error {
	_, err := client.Index().Index(config.EsIndexName).BodyJson(js).Do(context.Background())
	if err != nil {
		return err
	}
	// fmt.Printf("indexed %s to index %s, type %s \n", put.Id, put.Index, put.Type)
	return nil
}

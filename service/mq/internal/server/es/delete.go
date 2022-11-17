package es

import (
	"context"
	"essrv/service/mq/internal/config"
	"fmt"

	"github.com/olivere/elastic/v7"
)

func DeleteById(client *elastic.Client, id string) {
	_, err := client.Delete().
		Index(config.EsIndexName).
		Id(id).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
}

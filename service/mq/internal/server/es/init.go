package es

import (
	"context"
	"essrv/service/mq/internal/config"
	"fmt"
	"log"
	"os"

	"github.com/olivere/elastic/v7"
)

var client *elastic.Client

func InitEs() *elastic.Client {

	errorlog := log.New(os.Stdout, "app", log.LstdFlags)
	var err error
	client, err = elastic.NewClient(
		elastic.SetErrorLog(errorlog),
		elastic.SetURL(config.EsHost),
		elastic.SetSniff(false),
		elastic.SetBasicAuth("user", "secret"),
	)
	if err != nil {
		panic(err)
	}
	// info, code, err := client.Ping(config.EsHost).Do(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Es return with code %d and version %s \n", code, info.Version.Number)
	// esversionCode, err := client.ElasticsearchVersion(config.EsHost)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("es version %s\n", esversionCode)
	ctx := context.Background()

	exists, err := client.IndexExists(config.EsIndexName).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err := client.CreateIndex(config.EsIndexName).BodyString(config.EsMapping).Do(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println("新建Index:", config.EsIndexName)
	}
	return client
}

package es

import (
	"context"
	"essrv/service/mq/internal/config"
	"fmt"

	"github.com/olivere/elastic/v7"
)

// 根据id查询
func Get(client *elastic.Client, id string) {
	get, err := client.Get().Index(config.EsIndexName).Id(id).Do(context.Background())
	if err != nil {
		fmt.Println("err")
	}
	if get.Found {
		fmt.Printf("JOe got document %s in version %d from index %s,type %s \n", get.Id, get.Version, get.Index, get.Type)
		// get :::: &{device di h9wNcIQBPCIwOW-totdX    0x14000188e30 0x14000188e38 0x14000188e48 0x14000192600 true map[] <nil>}

	}
}

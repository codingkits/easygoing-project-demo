package es

import (
	"context"
	"essrv/service/mq/internal/config"
	"essrv/service/mq/internal/types"
	"fmt"
	"reflect"

	"github.com/olivere/elastic/v7"
)

// 插入之前查询并清理数据
func ClearUpData(client *elastic.Client, di types.DeviceInfo) error {
	bQ := elastic.NewBoolQuery().Must()
	t1 := elastic.NewTermQuery("idfa", di.Idfa)
	t2 := elastic.NewTermQuery("idfa_md5", di.IdfaMd5)
	t3 := elastic.NewTermQuery("imei", di.Imei)
	t4 := elastic.NewTermQuery("imei_md5", di.ImeiMd5)
	t5 := elastic.NewTermQuery("oaid", di.Oaid)
	t6 := elastic.NewTermQuery("oaid_md5", di.OaidMd5)
	bQ.Should(t1, t2, t3, t4, t5, t6)

	searchResult, err := client.Search().
		Index(config.EsIndexName).
		Query(bQ).
		Sort("ts", false).
		From(0).
		Size(10000).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		return err
	}
	// fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())

	if searchResult.TotalHits() > 0 {
		var ids []string
		for i := 0; i < int(searchResult.TotalHits()); i++ {
			ids = append(ids, searchResult.Hits.Hits[i].Id)
		}
		var newDi = types.DeviceInfo{}
		for _, item := range searchResult.Each(reflect.TypeOf(newDi)) {
			if t, ok := item.(types.DeviceInfo); ok {
				if t.Idfa != "" {
					newDi.Idfa = t.Idfa
				}
				if t.IdfaMd5 != "" {
					newDi.IdfaMd5 = t.IdfaMd5
				}
				if t.Imei != "" {
					newDi.Imei = t.Imei
				}
				if t.ImeiMd5 != "" {
					newDi.ImeiMd5 = t.ImeiMd5
				}
				if t.Oaid != "" {
					newDi.Oaid = t.Oaid
				}
				if t.OaidMd5 != "" {
					newDi.OaidMd5 = t.OaidMd5
				}
			}
		}
		e := StoreByStruct(client, newDi)
		if e != nil {
			return e
		}
		for n := 0; n < len(ids); n++ {
			DeleteById(client, ids[n])
		}
		fmt.Println("> 插入重新组装的完整数据:", newDi, ",并删除的ids:", ids)
	}

	fmt.Println("* 插入新数据:", di)
	return StoreByStruct(client, di)
}

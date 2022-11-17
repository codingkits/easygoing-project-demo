package logic

import (
	"essrv/service/mq/internal/server/es"

	"github.com/olivere/elastic/v7"
)

// type SyncDataToEsLogic struct {
// 	ctx    context.Context
// 	svcCtx *svc.ServiceContext
// }

var c *elastic.Client

func init() {
	c = es.InitEs()
}
func SyncDataToEsLogic(s string) error {
	di, ve := ValidJsonStr(s)
	if ve != nil {
		return ve
	}
	return es.ClearUpData(c, di)
}

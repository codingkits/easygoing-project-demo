package main

import (
	"github.com/gin-gonic/gin"
)

type DmpResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	r := gin.Default()
	r.GET("/hello", HelloAction)
	r.GET("/v1/dmp/query", DmpQueryAction)
	r.Run(":9001")
}

func HelloAction(ctx *gin.Context) {
	ctx.JSON(403, gin.H{
		"msg": "hello",
	})
}

func DmpQueryAction(ctx *gin.Context) {
	appKey := ctx.Query("app_key")
	if appKey == "joe" {
		ctx.JSON(200, gin.H{
			"msg": "success",
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg": "error",
		})
	}

}

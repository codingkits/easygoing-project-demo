package es

import (
	"encoding/json"
	"essrv/service/mq/internal/types"
	"fmt"
	"math/rand"
	"time"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(n int, acs ...[]rune) string {
	var letters []rune

	if len(acs) == 0 {
		letters = chars
	} else {
		letters = acs[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// 模拟消息
func MockMsg() string {
	m := map[string]string{
		"idfa":     "idfa_" + RandomString(8),
		"idfa_md5": "idfa_md5_" + RandomString(20),
		"imei":     "imei_" + RandomString(8),
		"imei_md5": "imei_md5_" + RandomString(20),
		"oaid":     "oaid_" + RandomString(8),
		"oaid_md5": "oaid_md5_" + RandomString(20),
		"ts":       string(time.Now().UnixMilli()),
	}
	mj, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
	}
	return string(mj)
}

func MockDi() types.DeviceInfo {
	return types.DeviceInfo{
		Idfa:    "idfa_" + RandomString(8),
		IdfaMd5: "idfa_md5_" + RandomString(20),
		Imei:    "imei_" + RandomString(8),
		ImeiMd5: "imei_md5_" + RandomString(20),
		Oaid:    "oaid_" + RandomString(8),
		OaidMd5: "oaid_md5_" + RandomString(20),
		Ts:      time.Now().UnixMilli(),
	}
}

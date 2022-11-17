package logic

import (
	"encoding/json"
	"errors"
	"essrv/service/mq/internal/types"
)

func ValidJsonStr(str string) (di types.DeviceInfo, err error) {
	err = json.Unmarshal([]byte(str), &di)
	if err != nil {
		return types.DeviceInfo{}, errors.New("invalid data")
	}
	if di == (types.DeviceInfo{}) {
		return types.DeviceInfo{}, errors.New("empty data")
	}
	return di, nil
}

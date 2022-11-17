package types

type DeviceInfo struct {
	Idfa    string `json:"idfa"`
	IdfaMd5 string `json:"idfa_md5"`
	Imei    string `json:"imei"`
	ImeiMd5 string `json:"imei_md5"`
	Oaid    string `json:"oaid"`
	OaidMd5 string `json:"oaid_md5"`
	Ts      int64  `json:"ts"`
}

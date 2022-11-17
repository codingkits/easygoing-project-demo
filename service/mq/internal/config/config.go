package config

const EsMapping = `
{
  "mappings": {
	"properties": {
		"idfa": {
			"type": "keyword"
		},
		"idfa_md5": {
			"type": "keyword"
		},
		"imei": {
			"type": "keyword"
		},
		"imei_md5": {
			"type": "keyword"
		},
		"oaid": {
			"type": "keyword"
		},
		"oaid_md5": {
			"type": "keyword"
		},
		"ts": {
			"type": "date"
		}
		}
  }
}`

var (
	SrvName     = "EsMqSrv"
	EsHost      = "http://localhost:9200"
	EsIndexName = "tzhorde"
	EsDocType   = "_doc"
	NsqD1Add    = "192.168.199.77:4150"
	NsqTopic    = "es_topic"
	NsqChan     = "es_chan"
)

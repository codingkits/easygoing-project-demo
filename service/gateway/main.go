// This plugin is used for dmp system auth.
package main

import (
	"bufio"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"
	// gopkg.in/yaml.v3
)

var HandlerRegisterer = registerer("krakend-auth-plugin")

var loger *log.Logger

var dmpConfig DmpConfig

var dmpJsonFileName = "xxx.json"

var dmpYamlFileName = "xxx.yaml"

type registerer string

type DmpConfig struct {
	Version       string          `yaml:"version" json:"version"`
	AppKeysConfig []AppKeysConfig `yaml:"app_keys" json:"app_keys"`
}

type AppKeysConfig struct {
	AppKey      string         `yaml:"app_key" json:"app_key"`
	AppSecret   string         `yaml:"app_secret" json:"app_secret"`
	Expires     string         `yaml:"expires" json:"expires"`
	Description string         `yaml:"description" json:"description"`
	IpPolicy    IpPolicyConfig `yaml:"ip_policy" json:"ip_policy"`
}

type IpPolicyConfig struct {
	Allow []string `yaml:"allow" json:"allow"`
	Deny  []string `yaml:"deny" json:"deny"`
}

func init() {
	logDirPath := "./kdlogs"
	isDirExist, _ := HasDir(logDirPath)
	if !isDirExist {
		err := os.Mkdir(logDirPath, os.ModePerm)
		if err != nil {
			fmt.Printf("create file err: %v\n", err)
		}
	}
	file := logDirPath + "/" + time.Now().Format("200601") + ".kd"
	logFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	loger = log.New(logFile, "[KrakenD Log] ", log.LstdFlags|log.Lshortfile)
	loger.Println("krakend-auth-plugin plugin loaded!")
	return
}

func (r registerer) RegisterHandlers(f func(
	name string,
	handler func(context.Context, map[string]interface{}, http.Handler) (http.Handler, error),
)) {
	f(string(r), r.registerHandlers)
}

func (r registerer) registerHandlers(_ context.Context, extra map[string]interface{}, h http.Handler) (http.Handler, error) {
	name, ok := extra["name"].([]interface{})

	if !ok {
		return nil, errors.New("config error")
	}

	if v, ok := name[0].(string); !ok || v != string(r) {
		return nil, fmt.Errorf("unknown register %s", name)
	}

	config, _ := extra["krakend-auth-plugin"].(map[string]interface{})

	// dmp_auth_config_yaml_file_cloud_path, _ := config["dmp_auth_config_yaml_file_cloud_path"].(string)
	// DownloadFileIntoLocal(dmp_auth_config_yaml_file_cloud_path, dmpYamlFileName)
	dmp_auth_config_json_file_cloud_path, _ := config["dmp_auth_config_json_file_cloud_path"].(string)
	DownloadFileIntoLocal(dmp_auth_config_json_file_cloud_path, dmpJsonFileName)

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		path, _ := config["path"].([]interface{})
		if !In(path, req.URL.Path) {
			loger.Println("[", req.URL.Path, "] is not in the hijack list,forward request right now!")
			h.ServeHTTP(w, req)
			return
		}

		loger.Println("[", req.URL.Path, "] in the hijack list,do the plugin logic and then  forward request!")
		appKey := req.URL.Query().Get("app_key")

		// dmpConfig := ReadLocalYamlFile(dmpYamlFileName)
		dmpConfig := ReadLocalJsonFile(dmpJsonFileName)

		secretMap := make(map[interface{}]interface{})
		for _, v := range dmpConfig.AppKeysConfig {
			secretMap[v.AppKey] = v.AppSecret
		}
		if secretMap[appKey] != nil {
			ts := req.URL.Query().Get("ts")
			nonce := req.URL.Query().Get("nonce")
			sign := req.URL.Query().Get("sign")
			if HmacSha256StrToHex(secretMap[appKey].(string), appKey, ts, nonce) == sign {
				loger.Println("Auth Success | current appKey:", appKey, " appSecret:", secretMap[appKey])
				h.ServeHTTP(w, req)
				return
			}
		} else {
			loger.Println("Auth Failure | current appKey:", appKey, " appSecret:", secretMap[appKey])
		}
		fmt.Fprintf(w, "Auth Failure")
	}), nil
}

func main() {}

func HmacSha256StrToHex(secret string, parts ...string) string {
	if len(parts) == 0 {
		return ""
	}
	h := hmac.New(sha256.New, []byte(secret))
	buf := bufio.NewWriter(h)
	for i := range parts {
		buf.WriteString(parts[i])
	}
	buf.Flush()
	return hex.EncodeToString(h.Sum(nil))
}

func HasDir(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// func ReadLocalYamlFile(path string) (dmpConfig DmpConfig) {
// 	content, _ := ioutil.ReadFile(path)
// 	_ = localYaml.Unmarshal(content, &dmpConfig)
// 	return dmpConfig
// }

func ReadLocalJsonFile(path string) (dmpConfig DmpConfig) {
	jsonFile, _ := os.Open(path)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &dmpConfig)
	return dmpConfig
}

func DownloadFileIntoLocal(path string, localFileName string) {
	loger.Println("Downlaod ", path, " Into Local : ", localFileName)
	resp, err := http.Get(path)
	if err != nil {
		fmt.Printf("http.get err: %v\n", err)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile(localFileName, data, 0644)
}

func In(haystack interface{}, needle interface{}) bool {
	sVal := reflect.ValueOf(haystack)
	kind := sVal.Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < sVal.Len(); i++ {
			if sVal.Index(i).Interface() == needle {
				return true
			}
		}
		return false
	}
	return false
}

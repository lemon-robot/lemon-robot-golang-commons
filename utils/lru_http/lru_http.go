package lru_http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"lemon-robot-golang-commons/logger"
	"net/http"
	"strings"
	"sync"
)

type LRUHttp struct {
	baseUrl      string
	commonHeader map[string]string
}

var instance *LRUHttp
var once sync.Once

func GetInstance() *LRUHttp {
	once.Do(func() {
		instance = &LRUHttp{
			commonHeader: make(map[string]string),
		}
	})
	return instance
}

func (i *LRUHttp) AppendCommonHeader(headerContent map[string]string) {
	for key := range headerContent {
		i.commonHeader[key] = headerContent[key]
	}
}

func (i *LRUHttp) RequestJson(method, reqUrl string, data interface{}, header map[string]string) (string, error) {
	url := reqUrl
	if strings.Index(reqUrl, "http") < 0 {
		url = i.baseUrl + reqUrl
	}
	jsonBytes, _ := json.Marshal(data)
	request, newRequestErr := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))
	if newRequestErr != nil {
		logger.Error("Can not create new request, occur a error", newRequestErr)
		return "", newRequestErr
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	for key := range i.commonHeader {
		request.Header.Set(key, i.commonHeader[key])
	}
	for key := range header {
		request.Header.Set(key, header[key])
	}
	client := &http.Client{}
	response, doRequestErr := client.Do(request)
	if doRequestErr != nil {
		logger.Error("Do request ["+method+"] occur a error: "+url, doRequestErr)
		return "", newRequestErr
	}
	defer response.Body.Close()
	body, readResponseErr := ioutil.ReadAll(response.Body)
	if doRequestErr != nil {
		logger.Error("Read response ["+url+"] occur a error", readResponseErr)
		return "", newRequestErr
	}
	return string(body), nil
}

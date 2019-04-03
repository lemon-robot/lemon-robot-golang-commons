package lruhttp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"lemon-robot-golang-commons/logger"
	"net/http"
	"strings"
)

var baseUrl = ""
var commonHeader map[string]string

func SetBaseUrl(url string) {
	baseUrl = url
}

func SetCommonHeader(header map[string]string) {
	commonHeader = header
}

func AppendCommonHeader(headerContent map[string]string) {
	if commonHeader == nil {
		commonHeader = headerContent
	} else {
		for key := range headerContent {
			commonHeader[key] = headerContent[key]
		}
	}
}

func RequestJson(method, reqUrl string, data interface{}, header map[string]string) (string, error) {
	url := reqUrl
	if strings.Index(reqUrl, "http") < 0 {
		url = baseUrl + reqUrl
	}
	jsonBytes, _ := json.Marshal(data)
	request, newRequestErr := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))
	if newRequestErr != nil {
		logger.Error("Can not create new request, occur a error", newRequestErr)
		return "", newRequestErr
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	for key := range commonHeader {
		request.Header.Set(key, commonHeader[key])
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

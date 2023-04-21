package go_rapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func HttpGet(reqUrl string, headers map[string]string) ([]byte, error) {
	return NewHttpRequest("GET", reqUrl, "", headers)
}

func HttpPostForm(reqUrl string, postData url.Values, headers map[string]string) ([]byte, error) {

	return NewHttpRequest("POST", reqUrl, postData.Encode(), headers)

}

func HttpPostBody(reqUrl string, postData map[string]interface{}, headers map[string]string) ([]byte, error) {
	if headers == nil {
		headers = map[string]string{}
	}

	headers["Content-Type"] = "application/json"
	data, _ := json.Marshal(postData)
	return NewHttpRequest("POST", reqUrl, string(data), headers)
}

func HttpPutBody(reqUrl string, postData map[string]interface{}, headers map[string]string) ([]byte, error) {
	if headers == nil {
		headers = map[string]string{}
	}

	headers["Content-Type"] = "application/json"
	data, _ := json.Marshal(postData)
	return NewHttpRequest("PUT", reqUrl, string(data), headers)
}

func NewHttpRequest(reqType string, reqUrl string, postData string, requstHeaders map[string]string) ([]byte, error) {
	fmt.Println(fmt.Sprintf("%s request url: %s", reqType, reqUrl))
	if postData != "" {
		fmt.Println(fmt.Sprintf("postData: %s", postData))
	}

	req, _ := http.NewRequest(reqType, reqUrl, strings.NewReader(postData))
	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36")
	}
	if requstHeaders != nil {
		for k, v := range requstHeaders {
			req.Header.Add(k, v)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 || resp.StatusCode <= 199 {
		return nil, errors.New(fmt.Sprintf("HttpStatusCode:%d ,Desc:%s", resp.StatusCode, string(bodyData)))
	}

	return bodyData, nil
}

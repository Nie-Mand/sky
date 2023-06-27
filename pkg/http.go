package pkg

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	baseUrl string
	headers map[string]string
}

var Http = HttpClient{}

func (c *HttpClient) Init(baseUrl string, headers map[string]string) {
	c.baseUrl = baseUrl
	c.headers = headers
}

func (c *HttpClient) Get(url string) (string, error) {
	resp, err := http.Get(c.baseUrl + url)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	return string(body), nil
}

func (c *HttpClient) Post(url string, data map[string]string) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	responseBody := bytes.NewBuffer(jsonData)
	resp, err := http.Post(c.baseUrl+url, "application/json", responseBody)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	return string(body), nil
}
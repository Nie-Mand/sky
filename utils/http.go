package utils

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

func (c *HttpClient) Download(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte(""), err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), nil
	}

	return body, nil
}

func (c *HttpClient) Post(url string, data map[string]interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	client := &http.Client{}

	responseBody := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest("POST", c.baseUrl+url, responseBody)

	if err != nil {
		return "", err
	}
	
	headers := c.headers
	headers["Content-Type"] = "application/json"

	
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	return string(body), nil
}
package http_util

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Read(urlPath string, timeout int, headers map[string]string) ([]byte, error) {
	// prepare http request
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest():%s", err.Error())
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// prepare http client
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			//DisableKeepAlives: true,
		},
	}

	// send request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client.do():%s", err.Error())
	}
	defer resp.Body.Close()

	// check status code
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code:%d", resp.StatusCode)
	}

	// read from response
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ReadAll():%s", err.Error())
	}
	return bytes, nil
}

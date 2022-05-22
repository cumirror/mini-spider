package http_util

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHttpRead(t *testing.T) {
	timeout := time.Second * 2
	header := map[string]string{
		"User": "mini-spider",
	}
	rspData := "OK"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/ok":
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, rspData)
		case "/timeout":
			time.Sleep(timeout)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, rspData)
		case "/header":
			if r.Header["User"][0] == "mini-spider" {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, rspData)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	}))
	defer svr.Close()

	// 1.test ok
	data, err := Read(svr.URL+"/ok", 0, nil)
	if err != nil || string(data) != rspData {
		t.Errorf("Read not as expect: %s %s", string(data), err)
	}

	// 2.1 test timeout: Within the threshold(2s)
	data, err = Read(svr.URL+"/timeout", 3, nil)
	if err != nil || string(data) != rspData {
		t.Errorf("Read not as expect: %s %s", string(data), err)
	}
	// 2.2 test timeout: Outside the threshold(2s)
	data, err = Read(svr.URL+"/timeout", 1, nil)
	if err == nil {
		t.Error("Read not as expect")
	}

	// 3.test header
	data, err = Read(svr.URL+"/header", 0, header)
	if err != nil || string(data) != rspData {
		t.Errorf("Read not as expect: %s %s", string(data), err)
	}
}

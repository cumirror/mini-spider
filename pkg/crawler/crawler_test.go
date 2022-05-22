/* crawler_test.go: test for crawler */
/*
modification history
--------------------
2017/07/21, by Xiongmin LIN, create
*/
/*
DESCRIPTION
*/

package crawler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mini_spider_config "github.com/cumirror/mini-spider/pkg/config"
	"github.com/cumirror/mini-spider/pkg/model"
)

func TestCrawler(t *testing.T) {
	html := `
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>mini-spider</title>
</head>
<body>
  <p>Absolute Path: <a href="https://just998.com/xianbao/41834973.html">a</a></p>
  <p>Relative Path: <a href="b.html">b</a></p>
</body>
</html>
`
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, html)
	}))
	defer svr.Close()

	urlTable := model.NewUrlTable()
	conf, _ := mini_spider_config.LoadConfig("../../test/spider.conf")

	var queue model.Queue
	queue.Init()
	queue.Add(&model.CrawlTask{svr.URL, 1, nil})

	c := NewCrawler(urlTable, &conf, &queue)

	go c.Run() // can't stop, so remove waitGroup
	time.Sleep(time.Second * 5)

	// check visit result
	verifyLinks := []string{
		"https://just998.com/xianbao/41834973.html", // use small site which would not get 403 error
		svr.URL + "/b.html",
	}

	for _, link := range verifyLinks {
		if !c.urlTable.Exist(link) {
			t.Errorf("%s not visited", link)
		}
	}
}

/* mini_spider.go */
/*
modification history
--------------------
2017/07/20, by Xiongmin LIN, create
*/
/*
DESCRIPTION
- mini spider
*/

package spider

import (
	mini_spider_config "github.com/cumirror/mini-spider/pkg/config"
	"github.com/cumirror/mini-spider/pkg/crawler"
	"github.com/cumirror/mini-spider/pkg/model"
)

type MiniSpider struct {
	config   *mini_spider_config.MiniSpiderConf
	urlTable *model.UrlTable
	queue    model.Queue
	crawlers []*crawler.Crawler
}

// create new mini-spider
func NewMiniSpider(conf *mini_spider_config.MiniSpiderConf, seeds []string) (*MiniSpider, error) {
	ms := new(MiniSpider)
	ms.config = conf

	// create url table
	ms.urlTable = model.NewUrlTable()

	// initialize queue
	ms.queue.Init()

	// add seeds to queue
	for _, seed := range seeds {
		task := &model.CrawlTask{Url: seed, Depth: 1, Header: make(map[string]string)}
		ms.queue.Add(task)
	}

	// create crawlers, thread count was defined in conf
	ms.crawlers = make([]*crawler.Crawler, 0)
	for i := 0; i < conf.Basic.ThreadCount; i++ {
		ms.crawlers = append(ms.crawlers, crawler.NewCrawler(ms.urlTable, ms.config, &ms.queue))
	}

	return ms, nil
}

// run mini spider
func (ms *MiniSpider) Run() {
	// start all crawlers
	for _, c := range ms.crawlers {
		go c.Run()
	}
}

// get number of unfinished task
func (ms *MiniSpider) GetUnfinished() int {
	return ms.queue.GetUnfinished()
}

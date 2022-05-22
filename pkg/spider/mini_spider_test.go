/* mini_spider_test.go: test for mini_spider.go */
/*
modification history
--------------------
2017/07/21, by Xiongmin LIN, create
*/
/*
DESCRIPTION
*/

package spider

import (
	"testing"

	mini_spider_config "github.com/cumirror/mini-spider/pkg/config"
	"github.com/cumirror/mini-spider/pkg/seed"
)

func TestNewMiniSpider(t *testing.T) {
	conf, _ := mini_spider_config.LoadConfig("./test/spider.conf")
	seeds, _ := seed.LoadSeedFile(conf.Basic.UrlListFile)
	_, err := NewMiniSpider(&conf, seeds)
	if err != nil {
		t.Errorf("err happen in NewMiniSpider:%s", err.Error())
	}
}

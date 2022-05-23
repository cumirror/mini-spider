/* webpage_save_test.go - test for webpage_save.go  */
/*
modification history
--------------------
2017/07/22, by Xiongmin LIN, create
*/
/*
DESCRIPTION
*/
package webpage

import (
	"testing"
)

// test for genFilePath()
func TestGenFilePath(t *testing.T) {
	rootPath := "../../test/webpage/output"
	url := "www.baidu.com"

	filePath := genFilePath(url, rootPath)
	if filePath != "../../test/webpage/output/www.baidu.com" {
		t.Errorf("err in genFilePath(), filePath=%s", filePath)
	}
}

// test for saveWebPage()
func TestSaveWebPage(t *testing.T) {
	rootPath := "../../test/webpage/output"
	url := "www.baidu.com"
	data := []byte("this is a test")

	err := SaveWebPage(rootPath, url, data)
	if err != nil {
		t.Errorf("err in saveWebPage(%s, %s):%s", rootPath, url, err.Error())
	}
}

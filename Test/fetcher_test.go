package Test

import (
	"fmt"
	"shSpider_plus/fetcher"
	"testing"
)

//命名有要求
func TestFetcher(t *testing.T) {
	//url:="http://bbs.vcb-s.com/"
	url := "http://album.zhenai.com/u/1791723931"
	//url:="https://www.bilibili.com/"
	bytes, e := fetcher.Fetcher(url)
	if e != nil {
		panic(e)
	}

	fmt.Printf("%s", bytes)
}

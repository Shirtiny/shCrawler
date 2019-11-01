package Test

import (
	"shSpider_plus/fetcher"
	"shSpider_plus/zhenai/parser"
	"testing"
)

func TestPaserProfile(t *testing.T) {
	url := "http://album.zhenai.com/u/1791723931"
	bytes, e := fetcher.Fetcher(url)
	if e != nil {
		panic(e)
	}
	parser.ParseProfile(bytes)
}

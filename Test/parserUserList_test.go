package Test

import (
	"fmt"
	"shSpider_plus/fetcher"
	"shSpider_plus/zhenai/parser"
	"testing"
)

func TestParseUserList(t *testing.T) {
	url := "http://www.zhenai.com/zhenghun/alashanmeng"
	bytes, e := fetcher.Fetcher(url)

	if e != nil {
		panic(e)
	}

	parseResult := parser.ParseUserList(bytes)

	fmt.Printf("%s", parseResult.Objects)

}

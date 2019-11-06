package vcbParser

import (
	"fmt"
	"shSpider_plus/fetcher"
	"testing"
)

func TestParseSection(t *testing.T) {
	bytes, e := fetcher.Fetcher("http://bbs.vcb-s.com")
	if e!=nil {
		println(e)
	}
	parseResult := ParseSectionList(bytes)

	fmt.Printf("%v+",parseResult)
}

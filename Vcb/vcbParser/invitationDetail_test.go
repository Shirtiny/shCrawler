package vcbParser

import (
	"fmt"
	"shSpider_plus/fetcher"
	"testing"
)

func TestParseInvitationDetail(t *testing.T) {
	bytes, e := fetcher.Fetcher("http://bbs.vcb-s.com/thread-1515-1-1.html")

	if e!=nil {
		fmt.Print(e)
		panic(e)
	}

	ParseInvitationDetail(bytes)
}

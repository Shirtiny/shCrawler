package vcbParser

import (
	"fmt"
	"shSpider_plus/fetcher"
	"testing"
)

func TestParseInvitationDetail(t *testing.T) {
	bytes, e := fetcher.Fetcher("http://bbs.vcb-s.com/thread-4915-1-3.html")

	if e!=nil {
		fmt.Print(e)
		panic(e)
	}

	ParseInvitationDetail(bytes,32,1515)
}

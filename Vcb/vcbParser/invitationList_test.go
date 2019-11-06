package vcbParser

import (
	"shSpider_plus/fetcher"
	"testing"
)

func TestParseInvitationList(t *testing.T) {

	bytes, e := fetcher.Fetcher("http://bbs.vcb-s.com/forum-38-1.html")

	if e!=nil {
		println(e)
		panic(e)
	}

	//fmt.Printf("%s",bytes)

	ParseInvitationList(bytes)

}

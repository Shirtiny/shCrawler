package vcbParser

import (
	"fmt"
	"shSpider_plus/fetcher"
	"testing"
)

func TestParseUserProfile(t *testing.T) {
//http://bbs.vcb-s.com/space-uid-928.html
//http://bbs.vcb-s.com/space-uid-7746.html
	bytes, err := fetcher.Fetcher("http://bbs.vcb-s.com/space-uid-13851.html")
	if err!=nil {
		fmt.Printf("%s \n",err)
		panic(err)
	}
	ParseUserProfile(bytes,122)
}

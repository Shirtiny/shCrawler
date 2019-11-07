package vcbParser

import (
	"fmt"
	"regexp"
	"shSpider_plus/engine"
)

//提取帖子标题 包括标签和title
var invitationTitleReg = regexp.MustCompile(`<a[^>]*>([^<]+)</a>\s*<span\s*id="thread_subject"\s*[^>]*>([^<]+)</span>`)

//用于提取帖子内容
var invitationContentReg = regexp.MustCompile(`<td\s*class="t_f"\s*id="postmessage_\d+">[\s\S]*?</td>`)

//帖子内容 用于将连接替换为http://bbs.vcb-s.com/开头
var contentLinkReg = regexp.MustCompile(`<(a|img)+?([^>]+)(src|href)+?="(?P<url>([^hjm]+?[^"]+))"`)

//帖子内容 用于显示图片文件资源
var imageFileReg = regexp.MustCompile(`<img+?([^>]+)src+?="([^"]+)"\s*zoomfile="(?P<imageFileUrl>([^"]+))"`)

//ParseInvitationDetail 解析帖子
func ParseInvitationDetail(bytes []byte) engine.ParseResult {
	//帖子分类、标题
	categorise, title := extractInvitationTitle(bytes)
	fmt.Printf("类别： %s，标题：%s", categorise,title)
	println()
	invitationContent := extarctInvitationContent(bytes)
	fmt.Printf("提取和替换后的帖子内容\n： %s", invitationContent)

	return engine.ParseResult{}
}

//提取和替换帖子内容
func extarctInvitationContent(bytes []byte) string {

	subMatch := invitationContentReg.FindSubmatch(bytes)
	var content []byte
	if len(subMatch) != 0 {
		content = subMatch[0]
	}
	if len(content)==0 {
		return ""
	}
	//替换为 ${n}表示引用匹配的下标为n的子串的内容
	link := "<${1}${2}${3}=\"http://bbs.vcb-s.com/${url}\""
	content = contentLinkReg.ReplaceAll(content, []byte(link))
	//替换图片文件资源
	imageFileUrlStr := "<img${1}src=\"http://bbs.vcb-s.com/${imageFileUrl}\""
	content = imageFileReg.ReplaceAll(content, []byte(imageFileUrlStr))
	//fmt.Printf("替换后： %s", content)

	return string(content)
}

//提取帖子标题 包括标签和title
func extractInvitationTitle(bytes []byte) (string, string) {
	submatch := invitationTitleReg.FindSubmatch(bytes)
	//fmt.Printf("categorise : %s title：%s", submatch[1], submatch[2])
	if len(submatch)==0 {
		return "",""
	}
	return string(submatch[1]), string(submatch[2])
}

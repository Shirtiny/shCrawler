package vcbParser

import (
	"fmt"
	"log"
	"regexp"
	"shSpider_plus/Vcb/vcbModel"
	"shSpider_plus/engine"
	"shSpider_plus/model"
	"strconv"
	"time"
)

//提取帖子标题 包括标签和title
var invitationTitleReg = regexp.MustCompile(`<a[^>]*>([^<]+)</a>\s*<span\s*id="thread_subject"\s*[^>]*>([^<]+)</span>`)

//提取帖子标题 只匹配title 因为有些帖子没有标签（主题）
var invitationTitleOnlyReg = regexp.MustCompile(`<span\s*id="thread_subject"\s*[^>]*>([^<]+)</span>`)

//查看数和回复数
var viewsAndReplyNumReg = regexp.MustCompile(`<span class="xg1">查看:</span> <span class="xi1">(\d+)</span><span\s*class="pipe">.*</span><span\s*class="xg1">回复:</span>\s*<span\s*class="xi1">(?P<replyNum>\d+)</span>`)

//帖子创建时间 字符串
var invitationCreatedTimeReg = regexp.MustCompile(`<em[^>]*>发表于\s*(\d+-\d+-\d+ \d+:\d+:\d+)\s*</em>`)

//作者id 从·只看该作者·的url中提取
var authorIdReg = regexp.MustCompile(`<a\s*.*href="[^"]+authorid=(\d+)"\s*[^>]*>\s*只看该作者\s*</a>`)

//用于提取帖子内容
var invitationContentReg = regexp.MustCompile(`<td\s*class="t_f"\s*id="postmessage_\d+">[\s\S]*?</td>`)

//用于提取帖子内容后续部分 因为有的帖子有 有的帖子没有 <div id="comment_30212" class="cm">
var invitationPattlReg = regexp.MustCompile(`(<div class="pattl">[\s\S]*?</div>)\s*<div\s*id="comment_\d+"\s*`)

//帖子内容 用于将连接替换为http://bbs.vcb-s.com/开头
var contentLinkReg = regexp.MustCompile(`<(a|img)+?([^>]+)(src|href)+?="(?P<url>([^hjm]+?[^"]+))"`)

//帖子内容 用于显示图片文件资源
var imageFileReg = regexp.MustCompile(`<img+?([^>]+)src+?="([^"]+)"\s*zoomfile="(?P<imageFileUrl>([^"]+))"`)

//帖子内容 匹配下载文件的连接 来替换
var dFileUrlReg = regexp.MustCompile(`<a+?\s*href+?="(.*forum\.php\?.*aid=([^&"]+)[^"]*)"[\s\S]*?target="_blank">([^<]+)</a>`)

//帖子内容 用于提取文件连接中的aid
var fileAidReg = regexp.MustCompile(`forum\.php\?.*aid=([^&"]+)&noupdate=yes&nothumb=yes`)

//ParseInvitationDetail 解析器，解析帖子
func ParseInvitationDetail(bytes []byte, sectionId int, invitationId int) engine.ParseResult {

	result := engine.ParseResult{}
	//版块id
	//fmt.Println("接收到的版块id为：", sectionId)

	//帖子id
	//fmt.Println("接收到的帖子id为：", invitationId)

	//作者id
	authorId := extractOneInt(authorIdReg, bytes)
	//fmt.Printf("作者id：%v \n",authorId)

	//用户url 交给用户解析器解析
	authorIdStr := strconv.Itoa(authorId)
	userUrl := "http://bbs.vcb-s.com/space-uid-" + authorIdStr + ".html"

	//创建时间
	createdTimeStr := extractOneString(invitationCreatedTimeReg, bytes)
	//fmt.Printf("创建时间：%s \n", createdTimeStr)

	//创建时间时间戳 毫秒数
	gmtCreated := time.Now().UnixNano() / 1e6
	//修改时间时间戳
	gmtModified := gmtCreated

	//回复数 查看数
	views, replyNum := extractViewsAndReplyNum(bytes)
	//fmt.Printf("查看%v ，回复%v ；\n", views, replyNum)

	//帖子分类、标题
	categorise, title := extractInvitationTitle(bytes)
	//fmt.Printf("类别： %s，标题：%s \n", categorise, title)

	//帖子内容
	invitationContent := extarctInvitationContent(bytes)
	//fmt.Printf("提取和替换后的帖子内容\n： %s", invitationContent)

	esModel := model.EsModel{
		Index: "vcbinvitation",
		Type:  "invitation",
		ID:    "",
	}
	invitation := vcbModel.Invitation{
		SectionId:    sectionId,
		InvitationId: invitationId,
		Categorise:   categorise,
		Title:        title,
		Content:      invitationContent,
		Views:        views,
		ReplyNum:     replyNum,
		CreatedTime:  createdTimeStr,
		GmtCreated:   gmtCreated,
		GmtModified:  gmtModified,
		AuthorId:     authorId,
	}
	esModel.Object = invitation
	result.Objects = []model.EsModel{esModel}

	//请求
	result.Requests = append(result.Requests, engine.Request{
		Url:        userUrl,
		ParserFunc: UserProfileParser(authorId),
	})

	return result
}

//提取和替换帖子内容
func extarctInvitationContent(bytes []byte) string {

	subMatch := invitationContentReg.FindSubmatch(bytes)
	var content []byte
	if len(subMatch) != 0 {
		content = subMatch[0]
	}
	if len(content) == 0 {
		return ""
	}

	pattlSub := invitationPattlReg.FindSubmatch(bytes)
	if len(pattlSub) != 0 {
		content = append(content, '\n')
		content = append(content, pattlSub[1]...)
	}
	//fmt.Printf("合并后的帖子内容： %s\n",content)

	//替换为 ${n}表示引用匹配的下标为n的子串的内容
	link := "<${1}${2}${3}=\"http://bbs.vcb-s.com/${url}\""
	content = contentLinkReg.ReplaceAll(content, []byte(link))

	//匹配文件url中的aid
	var aid string
	aid = ""
	imageUrlSubmatch := imageFileReg.FindAllSubmatch(content, -1)
	for _, urlSub := range imageUrlSubmatch {
		imageUrl := urlSub[3]
		aidMatchs := fileAidReg.FindSubmatch(imageUrl)
		aid = string(aidMatchs[1])
	}

	//替换图片文件资源 使用自站的shApi来代理处理图片资源 解决无法直接预览图片的问题
	imageFileUrlStr := "<img${1}src=\"/shApi/vcb/imageHelper?aid=" + aid + "\""
	content = imageFileReg.ReplaceAll(content, []byte(imageFileUrlStr))

	//替换文件下载连接
	dFileSub := dFileUrlReg.FindSubmatch(content)
	//fmt.Printf("%s\n %s\n %s\n %s\n", dFileSub[0], dFileSub[1], dFileSub[2], dFileSub[3])
	if len(dFileSub) != 0 {
		dFileUrlStr := "<a href=\"/shApi/vcb/dFileHelper?aid=${2}\" target=\"_blank\">${3}</a>"
		content = dFileUrlReg.ReplaceAll(content, []byte(dFileUrlStr))
	}

	fmt.Printf("最终文件内容： %s", content)

	return string(content)
}

//提取帖子标题 包括标签和title
func extractInvitationTitle(bytes []byte) (string, string) {
	submatch := invitationTitleReg.FindSubmatch(bytes)
	//fmt.Printf("categorise : %s title：%s", submatch[1], submatch[2])
	if len(submatch) == 0 {
		//如果主题+标题的模式没有匹配到，就尝试只匹配标题
		titleSubmatch := invitationTitleOnlyReg.FindSubmatch(bytes)
		if len(titleSubmatch) == 0 {
			return "", ""
		}
		return "", string(titleSubmatch[1])
	}
	return string(submatch[1]), string(submatch[2])
}

//提取查看数和回复数
func extractViewsAndReplyNum(bytes []byte) (int, int) {
	//匹配结果
	submatch := viewsAndReplyNumReg.FindSubmatch(bytes)

	if len(submatch) == 0 {
		return 0, 0
	} else {
		//字符串转int
		views, err1 := strconv.Atoi(string(submatch[1]))
		replyNum, err2 := strconv.Atoi(string(submatch[2]))
		//异常处理
		if err1 != nil || err2 != nil {
			log.Printf("帖子查看数或回复数字符串转int失败：%s，%s", err1, err2)
		}
		return views, replyNum
	}
}

//通用提取，只有一个子串（小括号） 返回字符串
func extractOneString(reg *regexp.Regexp, text []byte) string {
	submatch := reg.FindSubmatch(text)
	if len(submatch) == 0 {
		return ""
	}
	return string(submatch[1])
}

//通用提取，只有一个子串 返回int
func extractOneInt(reg *regexp.Regexp, text []byte) int {
	submatch := reg.FindSubmatch(text)
	if len(submatch) == 0 {
		return 0
	}
	i, e := strconv.Atoi(string(submatch[1]))
	if e != nil {
		log.Printf("字符串转换int失败")
		return 0
	}

	return i
}

//封装 为下一个用户解析器 传输用户id
func UserProfileParser(userId int) engine.ParserFunc {
	return func(bytes []byte) engine.ParseResult {
		return ParseUserProfile(bytes, userId)
	}
}

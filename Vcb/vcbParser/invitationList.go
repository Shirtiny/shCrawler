package vcbParser

import (
	"fmt"
	"log"
	"regexp"
	"shSpider_plus/engine"
	"strconv"
)

//版块
var sectionTitleReg = regexp.MustCompile(`<a[ \n]*href="forum.php\?gid=36"[ \n]*>[ \n]*论坛版块[ \n]*</a>[ \n]*<em>&rsaquo;[ \n]*</em>[ \n]*<a[ \n]*href="([^"]+)">([^<]+)</a>`)
//url和帖子标题 暂不关心帖子内容 只存储版块和帖子直接的联系
var invitationReg = regexp.MustCompile(`<a[ \n]*href="([^"]+)"[ \n\-();",:=#%!a-zA-Z0-9]*onclick="atarget\(this\)"[ \n]*class="s xst">([^<]+)</a>`)
//页码组
var pageNumsReg = regexp.MustCompile(`<a[ \n]*.*[ \n]*href=[ \n]*.*[ \n]*rel=[ \n]*.*[ \n]*curpage="([0-9]+)"[ \n]*.*[ \n]*totalpage="([0-9]+)"[ \n]*.*[ \n]*.*[ \n]*.*[ \n]*>[ \n]*.*[ \n]*下一页[ \n]*.*[ \n]*</a>`)

//帖子id thread-5031-1-2.html
var invitationIdReg = regexp.MustCompile(`thread-(\d+)-\d+-\d+\.html`)

//解析帖子列表
func ParseInvitationList(bytes []byte, sectionId int) engine.ParseResult {

	extract(sectionTitleReg, bytes, "版块名")

	invitations := extract(invitationReg, bytes, "帖子")

	pageNums := extract(pageNumsReg, bytes, "页码组")

	parseResult := engine.ParseResult{}

	//遍历帖子匹配结果 把每个URL放入新的request中
	limit := 0
	for _, invitation := range invitations {

		//提取id 输入thread-5031-1-2.html  返回5031
		InvitationId := extractInvitationId(invitation[0])

		//加入requests
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url: "http://bbs.vcb-s.com/" + string(invitation[0]),
			//向下一个解析器传入sectionId和invitationId
			ParserFunc: InvitationDetailParser(sectionId, InvitationId),
		})
	}

	//遍历页码组 找出下一页的地址
	for _, pageNum := range pageNums {
		//限制爬取的力度
		if limit == 1 {
			break
		}
		limit++

		//1是当前页 2是总页数
		//字符串转int
		curPage, err := strconv.Atoi(string(pageNum[0]))
		if err != nil {
			log.Printf("转换当前页码字符串到int时出错： %s", err)
		}
		totalPage, err := strconv.Atoi(string(pageNum[1]))
		if err != nil {
			log.Printf("转换总页数的字符串到int时出错： %s", err)
		}

		//当前页小于总页数时，把下一页加入请求队列
		nextPageStr := strconv.Itoa(curPage + 1)
		if curPage < totalPage {
			parseResult.Requests = append(parseResult.Requests, engine.Request{
				Url: "http://bbs.vcb-s.com/forum-37-" + nextPageStr + ".html",
				//解析器还是当前解析器
				ParserFunc: InvitationListParser(sectionId),
			})
		}
		fmt.Printf("当前页为：%s，总页数为：%s，下一页地址为：http://bbs.vcb-s.com/forum-37-"+nextPageStr+".html", pageNum[0], pageNum[1])
	}

	return parseResult
}

//提取 打印一下看看
func extract(reg *regexp.Regexp, bytes []byte, regName string) (matchs [][][]byte) {

	allSubmatch := reg.FindAllSubmatch(bytes, -1)

	for _, submatchs := range allSubmatch {
		//Url是submatchs[1]，帖子标题是submatchs[2]
		fmt.Printf("%s", submatchs[1:])
		match := submatchs[1:]
		matchs = append(matchs, match)
		fmt.Println()
	}
	fmt.Println("匹配到#", len(allSubmatch), "个"+regName)

	return matchs

}

//提取帖子id thread-5031-1-2.html  返回5031
func extractInvitationId(invitationNameBytes []byte) int {
	//正则匹配
	subMatch := invitationIdReg.FindSubmatch(invitationNameBytes)
	//字符串转int
	if len(subMatch) != 0 {

		invitationId, err := strconv.Atoi(string(subMatch[1]))
		//转换异常处理
		if err != nil {
			log.Printf("帖子id字符串转int失败：%s", err)
		}
		return invitationId

	} else {
		return 0
	}

}

//封装 向下一个解析器传递帖子id
func InvitationDetailParser(sectionId int, invitationId int) engine.ParserFunc {

	return func(bytes []byte) engine.ParseResult {
		return ParseInvitationDetail(bytes, sectionId, invitationId)
	}
}

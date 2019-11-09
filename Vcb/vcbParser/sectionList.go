package vcbParser

import (
	"fmt"
	"log"
	"regexp"
	"shSpider_plus/Vcb/vcbModel"
	"shSpider_plus/engine"
	"shSpider_plus/model"
	"strconv"
)

//所有版块 匹配版块图片、URL地址、标题、描述
var sectionReg = regexp.MustCompile(`<a[ \n]*href="[^"]+">[ \n]*<img[ \n]*src="([^"]+)"[ a-zA-Z."=\n]*/>[ \n]*</a>[ a-zA-Z.</>\n]*<h2><a href="([^"]+)">([^<]+)</a>.*</h2>[^<]*<p.*>([^<]+)</p>`)

//版块id 匹配如forum+-37-1.html 取37
var sectionIdReg = regexp.MustCompile(`[a-zA-Z]+-([0-9]+)-\d+\.html`)

//解析版块列表
func ParseSectionList(bytes []byte) engine.ParseResult {
	//正则表达式匹配结果
	allSubmatch := sectionReg.FindAllSubmatch(bytes, -1)

	//对象模型
	section := vcbModel.Section{}
	//es存储模型
	esModel := model.EsModel{}
	//es存储路径 index
	esModel.Index = "vcbsection"
	//es存储路径 type
	esModel.Type = "section"

	//es存储数组
	var esModels []model.EsModel
	//存储总解析结果
	result := engine.ParseResult{}

	limit :=0
	for _, submatchs := range allSubmatch {
		//限制爬取的力度
		if limit==1 {
			break
		}
		limit++

		//submatchs[1]是imageUrl，submatchs[2]是url，submatchs[3]是版块标题，submatchs[4]是版块描述
		fmt.Printf("%s %s %s %s", submatchs[1], submatchs[2], submatchs[3], submatchs[4])
		//存入模型
		section.ImageUrl = "http://bbs.vcb-s.com/" + string(submatchs[1])
		section.Title = string(submatchs[3])
		section.Description = string(submatchs[4])
		//提取版块id
		sectionId := extarctSectionId(submatchs[2])
		section.VcbSectionId = sectionId
		//es存储id 版块id 如forum-37-1 由forum-37-1.html截取得 需要自增的话，让id为空串
		esModel.ID = string(submatchs[2])[0:10]
		//存入es存储模型
		esModel.Object = section
		//存入es存储数组
		esModels = append(esModels, esModel)
		//解析到的需要再次处理的请求
		result.Requests = append(result.Requests, engine.Request{
			//请求地址
			Url: "http://bbs.vcb-s.com/" + string(submatchs[2]),
			//对应解析器
			ParserFunc: InvitationListParser(sectionId),
		})
		fmt.Println()
	}
	fmt.Println("匹配到#", len(allSubmatch), "个版块")

	//解析到的需要存储的对象
	result.Objects = esModels

	return result

}

//从forum-37-1.html中 提取37
func extarctSectionId(sectionNameBytes []byte) int {
	//正则表达式匹配
	subMatch := sectionIdReg.FindSubmatch(sectionNameBytes)
	//转为int
	if len(subMatch) != 0 {
		sectionId, err := strconv.Atoi(string(subMatch[1]))
		//异常处理
		if err != nil {
			log.Printf("版块id字符串转换出错 %s", err)
		}
		return sectionId
	} else {

		return 0
	}

}

//返回函数的封装 可以对下一个解析器传递参数
func InvitationListParser(sectionId int) engine.ParserFunc {
	return func(bytes []byte) engine.ParseResult {
		//在新建的匿名函数里，调用invitationList解析器 传入版块id
		return ParseInvitationList(bytes, sectionId)
	}
}

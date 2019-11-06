package vcbParser

import (
	"fmt"
	"regexp"
	"shSpider_plus/Vcb/vcbModel"
	"shSpider_plus/engine"
	"shSpider_plus/model"
)

//所有版块 匹配版块图片、URL地址、标题、描述
var sectionReg = regexp.MustCompile(`<a[ \n]*href="[^"]+">[ \n]*<img[ \n]*src="([^"]+)"[ a-zA-Z."=\n]*/>[ \n]*</a>[ a-zA-Z.</>\n]*<h2><a href="([^"]+)">([^<]+)</a>.*</h2>[^<]*<p.*>([^<]+)</p>`)
//版块id，暂时从版块url中抽取

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

	for _, submatchs := range allSubmatch {
		//submatchs[1]是imageUrl，submatchs[2]是url，submatchs[3]是版块标题，submatchs[4]是版块描述
		fmt.Printf("%s %s %s %s", submatchs[1], submatchs[2], submatchs[3], submatchs[4])
		//存入模型
		section.ImageUrl = "http://bbs.vcb-s.com/" + string(submatchs[1])
		section.Title = string(submatchs[3])
		section.Description = string(submatchs[4])
		section.VcbSectionId=string(submatchs[2])[0:10]
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
			ParserFunc: ParseInvitationList,
		})
		fmt.Println()
	}
	fmt.Println("匹配到#", len(allSubmatch), "个版块")

	//解析到的需要存储的对象
	result.Objects = esModels

	return result

}

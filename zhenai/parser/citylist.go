package parser

import (
	"fmt"
	"regexp"
	"shSpider_plus/engine"
)

//正则表达式 <a target="_blank" href="http://www.zhenai.com/zhenghun/beijing" data-v-7e67c21c>北京</a>
const cityReg = `<a [a-zA-Z0-9]*href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`

//解析url，返回城市列表以及对应请求,返回ParseResult
func ParseCityList(bytes []byte) engine.ParseResult {

	//编译正则表达式
	res := regexp.MustCompile(cityReg)

	//找出所有匹配到的子串
	matchStrings := res.FindAllSubmatch(bytes, -1)

	result := engine.ParseResult{}

	//限制一下分析次数，分析10个就行了
	//limit:=10
	//遍历[][]string
	for _, strArray := range matchStrings {
		//处理url到返回结果
		url := strArray[1]

		//把url存入结果中的Requests数组，每个request包含url和对应的解析方法
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(url),
			ParserFunc: ParseUserList,
		})

		//limit--
		//if limit==0 {
		//	break
		//}
	}
	fmt.Println("一共匹配到了：", len(matchStrings))

	return result
}

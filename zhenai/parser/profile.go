package parser

import (
	"regexp"
	"shSpider_plus/engine"
	"shSpider_plus/model"
	"strconv"
)

//编译好的正则表达式，省去每次都编译，节省性能

//id
var idReg = regexp.MustCompile(`<div class="id" [^>]*>ID：([\d]+)</div>`)

//昵称
var nickNameReg = regexp.MustCompile(`<h1 class="nickName" [^>]*>([^<]+)</h1>`)

//描述
var descriptionReg = regexp.MustCompile(`<div[^>]*>内心独白[^<]*</div>[^<]*<div[^>]*>[^<]*<span[^>]*>([^<]+)</span>.*</div>`)

//年龄
var ageReg = regexp.MustCompile(`<div[^>]*>([\d]+)岁</div>`)

//身高
var heightReg = regexp.MustCompile(`<div[^>]*>([\d]+)cm</div>`)

//体重
var weightReg = regexp.MustCompile(`<div[^>]*>([\d]+)kg</div>`)

//工作地
var locationReg = regexp.MustCompile(`<div[^>]*>工作地:([^<]+)</div>`)

//ParseProfile 解析个人资料信息
func ParseProfile(bytes []byte) engine.ParseResult {

	//用于存取profile信息的模型对象
	profile := model.Profile{}

	//id
	//调用提取数字的方法 string转int然后赋值
	profile.ID = extractIntWithReg(bytes, idReg)

	//昵称
	profile.NickName = extractStringWithReg(bytes, nickNameReg)

	//描述
	profile.Description = extractStringWithReg(bytes, descriptionReg)

	//年龄
	profile.Age = extractIntWithReg(bytes, ageReg)

	//身高
	profile.Height = extractIntWithReg(bytes, heightReg)

	//体重
	profile.Weight = extractIntWithReg(bytes, weightReg)

	//工作地
	profile.Location = extractStringWithReg(bytes, locationReg)

	return engine.ParseResult{
		Objects: []interface{}{profile},
	}
}

//使用正则表达式提取字符串 返回string
func extractStringWithReg(bytes []byte, reg *regexp.Regexp) string {
	//寻找子串
	submatch := reg.FindSubmatch(bytes)

	//如果正则表达式有匹配的结果，子串应该大于2个
	if len(submatch) >= 2 {
		//submatch[0]是完整的<h1 class="nickName" [^>]*>([^<]+)</h1>，而submatch[1]是()里的结果
		return string(submatch[1])

		//否则返回空串
	} else {
		return ""
	}
}

//提取数字 返回int
func extractIntWithReg(bytes []byte, reg *regexp.Regexp) int {
	//调用上面提取字符串的方法
	resStr := extractStringWithReg(bytes, reg)
	//把字符串转为int
	number, e := strconv.Atoi(resStr)
	//没出错返回数字，出错返回-1
	if e == nil {
		return number
	} else {
		return -1
	}
}

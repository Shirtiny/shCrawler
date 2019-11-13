package vcbParser

import (
	"fmt"
	"regexp"
	"shSpider_plus/Vcb/vcbModel"
	"shSpider_plus/engine"
	"shSpider_plus/model"
	"time"
)

//用户名 昵称
var userNameReg = regexp.MustCompile(`<h2\s*class="mbn">\s*([^<]+)<span\s*[^>]*>[^(]*\(UID: (\d+)\)[^<]*</span>\s*</h2>`)
//注册时间
var userCreatedTimeReg = regexp.MustCompile(`<li>\s*<em>\s*注册时间\s*</em>\s*([^<]+)\s*</li>`)
//头像
var userAvatarImageReg = regexp.MustCompile(`<a\s*href="space-uid-\d+\.html">\s*<img\s*src="(http://bbs\.vcb-s\.com/uc_server/avatar\.php\?[^"]+)"\s*/>\s*</a>`)

func ParseUserProfile(bytes []byte, userId int) (result engine.ParseResult) {
	//用户id
	id := int64(userId)
	//用户名 昵称
	nickName := extractUserName(bytes)
	//println(nickName)
	//注册时间
	createdTime := extractOneString(userCreatedTimeReg, bytes)
	//println(createdTime)
	//头像
	avatarImage := extractOneString(userAvatarImageReg, bytes)
	//println(avatarImage)
	//创建时间戳
	gmtCreated := time.Now().UnixNano() / 1e6
	//修改时间戳
	gmtModified := gmtCreated

	//用户模型
	user := vcbModel.User{
		UserId:      id,
		NickName:    nickName,
		PassWord:    "123456",
		AvatarImage: avatarImage,
		CreatedTime: createdTime,
		GmtCreate:   gmtCreated,
		GmtModified: gmtModified,
		Description: "大家好",
	}
	//es存储模型
	esModel := model.EsModel{
		Index: "vcbuser",
		Type:  "user",
		ID:    "",
	}
	esModel.Object=user
	//解析结果
	result.Objects = []model.EsModel{esModel}
	return result
}

//提取用户名 顺便打印uid
func extractUserName(bytes []byte) string {
	submatch := userNameReg.FindSubmatch(bytes)
	if len(submatch) != 0 {
		fmt.Printf("用户名为：%s，Uid为：%s", submatch[1], submatch[2])
		return string(submatch[1])
	} else {
		return ""
	}
}

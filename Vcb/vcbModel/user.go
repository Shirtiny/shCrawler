package vcbModel

type User struct {
	UserId int64 `gorm:"AUTO_INCREMENT;primary_key;not null;unique"`
	//昵称
	NickName string `gorm:"size:100"`
	PassWord string `gorm:"size:100"`
	Email string `gorm:"size:200"`
	//头像地址
	AvatarImage string `gorm:"size:500"`
	//github的id
	GithubId string `gorm:"size:500"`
	//创建时间 字符串 多的
	CreatedTime string
	//创建时间 时间戳 毫秒数 time.Now().UnixNano() / 1e6
	GmtCreate int64
	//修改时间 时间戳 毫秒数
	GmtModified int64
	//令牌
	Token string `gorm:"size:500"`
	//个人签名 描述
	Description string `gorm:"size:20"`
}

// 设置表名为`user`
func (User) TableName() string {
	return "user"
}
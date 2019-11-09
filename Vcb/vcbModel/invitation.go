package vcbModel

type Invitation struct {
	//版块id
	SectionId int `gorm:"column:section_id;type:bigint;unique_index:S_I_index"` // 创建唯一索引并命名，如果找到其他相同名称的索引则创建组合索引 S_I_index是索引的名字 自定义
	//在Vcb的帖子id
	InvitationId int `gorm:"column:invitation_id;AUTO_INCREMENT;type:bigint;primary_key;unique_index:S_I_index"` //sql语句为CREATE UNIQUE INDEX S_I_index ON `section_invitations`(vcb_section_id, invitation_id)引则创建组合索引 S_I_index是索引的名字 自定义
	//分类 多的
	Categorise string
	//帖子标题
	Title string
	//内容
	Content string `gorm:"size:10240000"`
	//查看数
	Views int `gorm:"column:views;type:bigint;"`
	//回复数
	ReplyNum int `gorm:"column:reply_num;type:bigint;"`
	//创建时间 字符串 多的
	CreatedTime string
	//创建时间 时间戳 毫秒数 time.Now().UnixNano() / 1e6
	GmtCreated int64 `gorm:"column:gmt_created;"`
	//修改时间 时间戳 毫秒数
	GmtModified int64 `gorm:"column:gmt_modified;"`
	//发帖者id
	AuthorId int `gorm:"column:author_id;type:bigint;"`
	//糖果数
	CandyNum int `gorm:"column:candy_num;type:bigint;"`
}

// 设置表名为`invitation`
func (Invitation) TableName() string {
	return "invitation"
}

package vcbModel

type Invitation struct {
	//版块id
	SectionId int `gorm:"unique_index:S_I_index"` // 创建唯一索引并命名，如果找到其他相同名称的索引则创建组合索引 S_I_index是索引的名字 自定义
	//在Vcb的帖子id
	InvitationId int `gorm:"unique_index:S_I_index"` //sql语句为CREATE UNIQUE INDEX S_I_index ON `section_invitations`(vcb_section_id, invitation_id)引则创建组合索引 S_I_index是索引的名字 自定义
	//分类
	Categorise string
	//帖子标题
	Title string
	//内容
	Content string `gorm:"size:10240000"`
	//查看数
	Views int
	//回复数
	ReplyNum int
	//创建时间
	CreatedTime string
	//发帖者id
	AuthorId int
}

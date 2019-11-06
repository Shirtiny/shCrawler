package vcbModel

type Invitation struct {
	//在Vcb的帖子id 唯一
	InvitationId string `gorm:"not null;unique"`
	//帖子标题
	Title string
	//内容
	Content string
	//发帖者id
}

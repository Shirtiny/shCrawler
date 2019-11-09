package vcbModel

type Section struct {
	//在Vcb的板块id 唯一
	VcbSectionId int `gorm:"column:section_id;type:bigint;not null;unique"`
	//版块标题
	Title string `gorm:"column:section_title"`
	//描述
	Description string `gorm:"column:section_description"`
	//logo图片地址
	ImageUrl string `gorm:"column:section_avatar_image;size:1000"`
	//版主id
	SectionOwnerId int `gorm:"column:section_owner_id;type:bigint;"`
	//帖子总数
	InvitationTotalNum int `gorm:"column:invitation_total_num;type:bigint;"`
	//版块总糖数
	SectionTotalCandy int `gorm:"column:section_total_candy;type:bigint;"`
	//板块评分
	SectionRate float64 `gotm:"column:section_rate;type:decimal;"`
}

// 设置Section表名为`section`
func (Section) TableName() string {
	return "section"
}
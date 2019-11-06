package vcbModel


type Section struct {
	//在Vcb的板块id 唯一
	VcbSectionId string `gorm:"not null;unique"`
	//版块标题
	Title string
	//描述
	Description string
	//logo图片地址
	ImageUrl string `gorm:"size:1000"`
}

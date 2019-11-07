package vcbModel

//中间表 关联版块和帖子
type SectionInvitation struct {

	VcbSectionId int `gorm:"unique_index:idx_name_code"` // 创建唯一索引并命名，如果找到其他相同名称的索引则创建组合索引 idx_name_code是索引的名字 自定义

	InvitationId int `gorm:"unique_index:idx_name_code"` //sql语句为CREATE UNIQUE INDEX idx_name_code ON `section_invitations`(vcb_section_id, invitation_id)
}

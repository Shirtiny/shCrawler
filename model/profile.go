package model

//EsModel 通用模型（可以在存储时读取id和type）
type EsModel struct {
	Index string
	Type string
	ID string
	//任意模型
	Object interface{}
 }

type Profile struct {
	//昵称
	NickName string
	//描述
	Description string
	//年龄
	Age int
	//身高
	Height int
	//体重
	Weight int
	//地理位置
	Location string
}


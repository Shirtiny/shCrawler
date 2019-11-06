package dataBase

import (
	"fmt"
	"shSpider_plus/Vcb/vcbModel"
	"testing"
)

func TestConnect(t *testing.T) {
	Connect()

	section := vcbModel.Section{
		Title:       "数据库",
		Description: "哈哈",
		ImageUrl:    "吃法",
	}

	e := DB.Create(section).Error
	if e!=nil {
		fmt.Printf("%s",e)
	}
}

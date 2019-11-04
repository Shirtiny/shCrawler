package Test

import (
	"fmt"
	"shSpider_plus/elastic"
	"shSpider_plus/model"
	"testing"
)

func TestEsCrud(t *testing.T) {
	client, e := elastic.NewClient("http://47.75.138.227:9200")

	if e != nil {
		fmt.Println("es客户端创建失败", e)
		panic(e)
	}

	esMode := model.EsModel{}
	esMode.Index="zhenaidata"
	esMode.Type="zhenai"
	esMode.ID="789"

	profile := model.Profile{
		NickName:    "shirtinyES2",
		Description: "go客户端创建的",
	}

	esMode.Object=profile

	elastic.Add(client, esMode)

	res := elastic.Search(client, "zhenaidata","zhenai","", "",20)

	fmt.Printf("%v+", res)
}

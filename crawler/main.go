package main

import (
	"fmt"
	"shSpider_plus"
	"shSpider_plus/dataBase"
	"shSpider_plus/elastic"
)

func main() {

	//连接数据库 为变量DB赋值
	dataBase.Connect()

	//创建es客户端
	url := "http://47.75.163.92:9200"
	client, e := elastic.NewClient(url)

	if e != nil {
		fmt.Println("es客户端创建失败", e)
		//panic(e)
	}

	//存储路径 解析时传入 每个不同的解析器有不同的存储路径
	//index := "zhenaidata"
	//index := "vcbdata"
	//esType := "zhenai"
	//esType := "vcb"

	//启动珍爱网的爬虫
	//shSpider_plus.StartZhenAi(client)

	//启动vcb的爬虫
	shSpider_plus.StartVcb(client)
}

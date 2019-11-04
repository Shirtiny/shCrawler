package main

import (
	"fmt"
	"shSpider_plus/elastic"
	"shSpider_plus/engine"
	"shSpider_plus/persist"
	"shSpider_plus/scheduler"
	"shSpider_plus/zhenai/parser"
)

func main() {
	//创建es客户端
	url := "http://47.75.138.227:9200"
	client, e := elastic.NewClient(url)

	if e != nil {
		fmt.Println("es客户端创建失败", e)
		panic(e)
	}

	//存储路径
	index := "zhenaidata"
	esType := "zhenai"

	//完成fetcher、parser、scheduler、engine es存储
	engineMain := engine.ConcurrentQueue{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 50,
		SaverChan:   persist.Saver(client,index ,esType),
	}

	engineMain.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}

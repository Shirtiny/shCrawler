package main

import (
	"shSpider_plus/engine"
	"shSpider_plus/persist"
	"shSpider_plus/scheduler"
	"shSpider_plus/zhenai/parser"
)

func main() {
//完成fetcher、parser、scheduler、engine，暂时没写存储
	engineMain := engine.ConcurrentQueue{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 50,
		SaverChan:   persist.Saver(),
	}

	engineMain.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

package shSpider_plus

import (
	"github.com/elastic/go-elasticsearch/v7"
	"shSpider_plus/Vcb/vcbParser"
	"shSpider_plus/engine"
	"shSpider_plus/persist"
	"shSpider_plus/scheduler"
	"shSpider_plus/zhenai/parser"
)

//StartZhenAi 启动珍爱网的爬虫
func StartZhenAi(client *elasticsearch.Client)  {
	//完成fetcher、parser、scheduler、engine es存储
	engineMain := engine.ConcurrentQueue{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 50,
		SaverChan:   persist.Saver(client),
	}

	engineMain.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

//启动Vcb的爬虫
func StartVcb(client *elasticsearch.Client)  {
	engineMain := engine.ConcurrentQueue{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 9,
		SaverChan:   persist.Saver(client),
	}

	//从版块列表开始
	engineMain.Run(engine.Request{
		Url:        "http://bbs.vcb-s.com",
		ParserFunc: vcbParser.ParseSectionList,
	})

	//从某个版块开始
	//engineMain.Run(engine.Request{
	//	Url:        "http://bbs.vcb-s.com/forum-37-1.html",
	//	ParserFunc: vcbParser.ParseInvitationList,
	//})
}
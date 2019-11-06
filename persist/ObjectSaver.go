package persist

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"shSpider_plus/dataBase"
	"shSpider_plus/model"
)

func Saver(client *elasticsearch.Client) chan model.EsModel {

	//输出通道
	out := make(chan model.EsModel)
	go func(_ *elasticsearch.Client) {
		count := 0
		for {
			object := <- out
			//存入es
			//elastic.Add(client, object)
			log.Printf("存储管道输出：#%d %v\n", count, object)
			//存入数据库 这里不能取地址
			dataBase.DB.Create(object.Object)
			count++
		}
	}(client)
	return out
}

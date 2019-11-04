package persist

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"shSpider_plus/elastic"
	"shSpider_plus/model"
)

func Saver(client *elasticsearch.Client,index string,esType string) chan model.EsModel {

	//输出通道
	out := make(chan model.EsModel)
	go func(client *elasticsearch.Client) {
		count := 0
		for {
			object := <-out
			elastic.Add(client,index,esType,object)
			log.Printf("worker输出的对象为：#%d %v\n", count, object)
			count++
		}
	}(client)
	return out
}

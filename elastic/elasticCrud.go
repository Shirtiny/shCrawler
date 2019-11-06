package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"shSpider_plus/model"
	"strings"
)

var jsonString string
//var esUrl="http://47.75.138.227:9200"
//var Size=20

//NewClient 创建客户端
func NewClient(url string) (*elasticsearch.Client,error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			url,
		},
	}
	return elasticsearch.NewClient(cfg)
}

//Add 增加和修改数据 使用传入EsModel对象的index和type
func Add(es *elasticsearch.Client,object model.EsModel) *esapi.Response {
	//对象转json
	bytes2, _ := json.Marshal(object)
	jsonString=string(bytes2)

	//增 改 index/type/id body
	request := esapi.IndexRequest{
		Index:        object.Index,
		DocumentType: object.Type,
		Body:         strings.NewReader(jsonString),
		Refresh:      "true",
	}

	//id不为空串时，才使用传入的id
	if object.ID!=""{
		request.DocumentID=object.ID
	}

	//插入
	response, err := request.Do(context.Background(), es)

	if err != nil {
		log.Fatalf("es插入或更新出错: %s", err)
	}

	defer response.Body.Close()

	fmt.Printf("插入结果 %s\n",response)
	return response

}

//搜索
func Search(es *elasticsearch.Client,esIndex string,esType string,fieldName string,fieldValue string,size int)  map[string]interface{}{
	var buf bytes.Buffer

	//查询map集合
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				fieldName: fieldValue,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("query json编码失败: %s", err)
	}

	//执行搜索
	response, e := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(esIndex),
		//es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
		es.Search.WithSize(size),
		es.Search.WithDocumentType(esType),
	)

	if e != nil {
		log.Fatalf("es搜索无响应: %s", e)
	}

	defer response.Body.Close()

	//map集合
	var res map[string]interface{}
	err := json.NewDecoder(response.Body).Decode(&res)

	if err != nil {
		log.Fatalf("responseBody解析失败: %s", err)
	}

	//返回结果map集合
	return res
}
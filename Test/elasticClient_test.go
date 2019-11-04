package Test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strings"
	"testing"
)

func TestElastic(t *testing.T) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://47.75.138.227:9200/",
		},
	}

	es, _ := elasticsearch.NewClient(cfg)

	//增 改 index/type/id body
	request := esapi.IndexRequest{
		Index:        "database",
		DocumentType: "user",
		DocumentID:   "999",
		Body:         strings.NewReader(`{"name":"client"}`),
		Refresh:      "true",
	}

	response, e := request.Do(context.Background(), es)

	if e != nil {
		panic(e)
	}

	defer response.Body.Close()

	log.Println(response)

	//查
	var buf bytes.Buffer

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"name": "shirtiny",
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	response, e = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("database"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)

	if e != nil {
		log.Fatalf("Error getting response: %s", e)
	}

	defer response.Body.Close()

	//map集合
	var r map[string]interface{}
	err := json.NewDecoder(response.Body).Decode(&r)
	if err != nil {
		log.Fatalf("responseBody解析失败: %s", err)
	}
	fmt.Printf("%v+", r)
}

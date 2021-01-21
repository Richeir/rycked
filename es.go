package rycked

import (
	"context"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

// EsBridge 1
type EsBridge struct {
}

// GetClient 1
func GetClient() *elasticsearch.Client {
	//TODO:到时候改成读 yaml 文件中的配置
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://192.168.2.163:9200",
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	return es
}

// WriteEs 1
func WriteEs(jsonString string, docID string) {
	esClient := GetClient()

	req := esapi.IndexRequest{
		Index:      "my-index",
		DocumentID: docID,
		Body:       strings.NewReader(jsonString),
	}
	req.Do(context.Background(), esClient)
}

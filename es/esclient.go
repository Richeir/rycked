package es

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

// TracerIndexName 1
const TracerIndexName = "rycked.tracer"

// SpanIndexName 1
const SpanIndexName = "rycked.span"

// EsClient 1
type EsClient struct {
}

// getClient 1
func getClient() *elasticsearch.Client {
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
func WriteEs(indexName, jsonString, docID string) {
	esClient := getClient()

	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: docID,
		Body:       strings.NewReader(jsonString),
	}
	req.Do(context.Background(), esClient)
}

// QueryTracer 1
func QueryTracer(tracerId string) *esapi.Response {
	es := getClient()

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"ID": tracerId,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(TracerIndexName),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error query es: %s", err)
	}
	//defer res.Body.Close()

	return res
}

// QueryTracer 1
func QuerySpan(spanId string) *esapi.Response {
	es := getClient()

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"ID": spanId,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(SpanIndexName),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error query es: %s", err)
	}
	//defer res.Body.Close()

	return res
}

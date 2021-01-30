package rycked

import (
	"encoding/json"
	"github.com/Richeir/rycked/es"
	"github.com/opentracing/opentracing-go"
	"log"
	"time"
)

// SpanContext 1
type SpanContext struct {
	TraceID string
	SpanId  string
	Baggage map[string]string
}

// ForeachBaggageItem 1
// TODO: 实现ForeachBaggageItem
func (sc *SpanContext) ForeachBaggageItem(handler func(k, v string) bool) {

}

// Span 一个有开始和结束的最小调用的过程
type Span struct {
	DocumentID    string
	ID            string
	TraceID       string
	OperationName string
	Depth         int
	StartAt       time.Time
	FinishAt      time.Time
	Tags          []opentracing.Tag
	Logs          []opentracing.LogData
}

func (span *Span) SpanContext() *SpanContext {
	sc := &SpanContext{}
	sc.TraceID = span.TraceID
	sc.SpanId = span.ID

	return sc
}

func (span *Span) SetTag(key string, value interface{}) Span {
	var tag opentracing.Tag
	tag.Key = key
	tag.Value = value
	span.Tags = append(span.Tags,tag)
	return *span
}

// Finish : Span 的完成方法，标记好结束时间，更新至ES
func (span *Span) Finish() {
	span.FinishAt = time.Now().UTC()
	b, _ := json.Marshal(span)
	es.WriteEs(es.SpanIndexName, string(b), span.DocumentID)
}

// GetSpan 直接返回Span对象
func GetSpan(spanid string) *Span {
	var (
		r map[string]interface{}
	)

	var res = es.QuerySpan(spanid)
	span := &Span{}

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	} else {
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		}

		log.Printf(
			"[%s] %d hits; took: %dms",
			res.Status(),
			int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
			int(r["took"].(float64)),
		)

		log.Print("hits content:")
		hitContent := r["hits"].(map[string]interface{})["hits"].([]interface{})

		if len(hitContent) > 0 {
			layout := "2006-01-02T15:04:05.0000000Z"
			var tracerObj = hitContent[0].(map[string]interface{})["_source"]
			span.ID = tracerObj.(map[string]interface{})["ID"].(string)
			span.DocumentID = tracerObj.(map[string]interface{})["DocumentID"].(string)
			span.TraceID = tracerObj.(map[string]interface{})["TraceID"].(string)
			span.Depth = int(tracerObj.(map[string]interface{})["Depth"].(float64))
			span.OperationName = tracerObj.(map[string]interface{})["OperationName"].(string)
			//span.Name = tracerObj.(map[string]interface{})["Name"].(string)
			span.StartAt, _ = time.Parse(layout, tracerObj.(map[string]interface{})["StartAt"].(string))
			span.FinishAt, _ = time.Parse(layout, tracerObj.(map[string]interface{})["FinishAt"].(string))
		}
	}

	return span
}

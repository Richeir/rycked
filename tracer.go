package rycked

import (
	"encoding/json"
	"github.com/Richeir/rycked/es"
	"log"
	"time"

	"github.com/google/uuid"
)

// Tracer 1
type Tracer struct {
	DocumentID string
	ID         string
	Name       string
	StartAt    time.Time
	FinishAt   time.Time
}

// SpanReferenceType is an enum type describing different categories of
// relationships between two Spans. If Span-2 refers to Span-1, the
// SpanReferenceType describes Span-1 from Span-2's perspective. For example,
// ChildOfRef means that Span-1 created Span-2.
//
// NOTE: Span-1 and Span-2 do *not* necessarily depend on each other for
// completion; e.g., Span-2 may be part of a background job enqueued by Span-1,
// or Span-2 may be sitting in a distributed queue behind Span-1.
type SpanReferenceType int

// NewTracer 创建新的 Tracer
func NewTracer(serviceName string) *Tracer {
	uuid, _ := uuid.NewUUID()
	tracer := &Tracer{
		ID:         uuid.String(),
		Name:       serviceName,
		StartAt:    time.Now().UTC(),
		DocumentID: uuid.String(),
	}

	b, _ := json.Marshal(tracer)
	es.WriteEs(es.TracerIndexName, string(b), uuid.String())
	return tracer
}

// NewSpan 创建 Span
func NewSpan(tracer *Tracer, operationName string) *Span {
	uuid, _ := uuid.NewUUID()
	span := &Span{
		ID:            uuid.String(),
		TraceID:       tracer.ID,
		OperationName: operationName,
		Depth:         1,
		StartAt:       time.Now().UTC(),
		DocumentID:    uuid.String(),
	}

	b, _ := json.Marshal(span)
	es.WriteEs(es.SpanIndexName, string(b), uuid.String())

	return span
}

// GetTracer 直接返回Tracer对象
func GetTracer(tracerid string) *Tracer {
	var (
		r map[string]interface{}
	)

	var res = es.QueryTracer(tracerid)
	tracer := &Tracer{}

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
		//log.Print(hitContent)
		//log.Print(hitContent[0])
		//log.Print(hitContent[0].(map[string]interface{}))
		//log.Print(hitContent[0].(map[string]interface{})["_source"])
		//log.Print(hitContent[0].(map[string]interface{})["_source"].(map[string]interface{})["Name"])

		if len(hitContent) > 0 {
			layout := "2006-01-02T15:04:05.0000000Z"
			var tracerObj = hitContent[0].(map[string]interface{})["_source"]
			tracer.ID = tracerObj.(map[string]interface{})["ID"].(string)
			tracer.DocumentID = tracerObj.(map[string]interface{})["DocumentID"].(string)
			tracer.Name = tracerObj.(map[string]interface{})["Name"].(string)
			tracer.StartAt, _ = time.Parse(layout, tracerObj.(map[string]interface{})["StartAt"].(string))
			tracer.FinishAt, _ = time.Parse(layout, tracerObj.(map[string]interface{})["FinishAt"].(string))
			//t.Log(tracerObj.(map[string]interface{})["StartAt"].(string))
			//
			//str := "2021-01-24T00:22:33.9831968+08:00"
			//t, err := time.Parse(layout, str)
			//if err != nil {
			//	log.Print(err)
			//}
			//log.Print(t)
		}

		//for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		//	log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
		//}
	}

	return tracer
}

// StartSpan 开启新的 span
// 如果传入的是 tracerid 可以找到对应的 tracer，那么返回，否则创建新的
func (t *Tracer) StartSpan(operationName string, tracerid string) *Span {
	var tracer *Tracer

	if tracerid == "" {
		tracerNew := NewTracer(operationName)
		tracer = tracerNew
	} else {
		//读取 Tracer
	}

	newSpan := NewSpan(tracer, operationName)
	return newSpan
}

//
//func (t *Tracer) GetTracer(tracerid string) *Tracer{
//
//}

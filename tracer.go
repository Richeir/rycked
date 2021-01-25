package apm

import (
	"encoding/json"
	"time"

	"github.com/elastic/go-elasticsearch/esapi"

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
		ID:      uuid.String(),
		Name:    serviceName,
		StartAt: time.Now(),
	}

	b, _ := json.Marshal(tracer)
	WriteEs(TracerIndexName, string(b), "")
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
		StartAt:       time.Now(),
	}

	//TODO:写入ES
	b, _ := json.Marshal(span)
	WriteEs(SpanIndexName, string(b), "")

	return span
}

// GetTracer 1
func GetTracer(tracerid string) *esapi.Response {
	return QueryTracer(tracerid)
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

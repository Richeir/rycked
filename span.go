package rycked

import (
	"encoding/json"
	"time"
)

//
type SpanContext interface {
}

//一个有开始和结束的最小调用的过程
type Span struct {
	DocumentID string
	Id            string
	TraceId       string
	OperationName string
	Depth         int
	StartAt       int64
	FinishAt      int64
}

type SpanReference struct {
	Type              SpanReferenceType
	ReferencedContext SpanContext
}

//Span 的完成方法，标记好结束时间，更新至ES
func (span *Span) Finish(targetSpan *Span) {
	targetSpan.FinishAt = time.Now().UnixNano() / 1e6
	//TODO: 更新至ES
	b, _ := json.Marshal(span)
	WriteEs(string(b), targetSpan.DocumentID)
}
package rycked

import (
	"encoding/json"
	"time"
)

// SpanContext 1
type SpanContext interface {
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
}

// Finish : Span 的完成方法，标记好结束时间，更新至ES
func (span *Span) Finish(targetSpan *Span) {
	targetSpan.FinishAt = time.Now()
	//TODO: 更新至ES
	b, _ := json.Marshal(span)
	WriteEs(SpanIndexName, string(b), targetSpan.DocumentID)
}

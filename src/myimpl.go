package richeir

type myTracer struct {
}

func (a myTracer) ForeachBaggageItem(handler func(k, v string) bool) {

}

// func (a opentracing.Span) FinishAt() {

// }

// MySpan is my tracer
type MySpan struct {
	a int
}

func (r MySpan) Finish() {

}

func (r MySpan) Context() SpanContext {
	return nil
}

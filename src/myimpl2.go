package richeir
//
//// SpanContext represents Span state that must propagate to descendant Spans and across process
//// boundaries (e.g., a <trace_id, span_id, sampled> tuple).
//type SpanContext interface {
//	// ForeachBaggageItem grants access to all baggage items stored in the
//	// SpanContext.
//	// The handler function will be called for each baggage key/value pair.
//	// The ordering of items is not guaranteed.
//	//
//	// The bool return value indicates if the handler wants to continue iterating
//	// through the rest of the baggage items; for example if the handler is trying to
//	// find some baggage item by pattern matching the name, it can return false
//	// as soon as the item is found to stop further iterations.
//	ForeachBaggageItem(handler func(k, v string) bool)
//}
//
//type dong struct {
//}
//
//func (d dong) ForeachBaggageItem(handler func(k, v string) bool) {
//
//}

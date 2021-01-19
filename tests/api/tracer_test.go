package api

import (
	"github.com/Richeir/rycked"
	"testing"
)

func TestTracer(t *testing.T) {
	var tracer rycked.Tracer
	tracer.StartSpan("n1","")
}

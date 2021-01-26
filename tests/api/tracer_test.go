package api

import (
	"testing"

	"github.com/Richeir/rycked"
)

func TestTracer(t *testing.T) {
	var tracer rycked.Tracer
	tracer.StartSpan("testTracer1", "")
}

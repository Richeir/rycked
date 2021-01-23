package api

import (
	"github.com/Richeir/rycked"
	"testing"
)

func TestTracer(t *testing.T) {
	var tracer apm.Tracer
	tracer.StartSpan("testTracer1", "")
}

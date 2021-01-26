package api

import (
	"testing"

	"github.com/Richeir/rycked"
)

func TestCreateTracer(t *testing.T) {
	var tracer rycked.Tracer
	tracer.StartSpan("testTracer1", "")
}

func TestGetTracer(t *testing.T)  {
	tracer1 := rycked.GetTracer("fc560e8c-5ff8-11eb-b426-6c626debff99")
	t.Log((tracer1))
}

func TestGetSpan(t *testing.T)  {
	span1 := rycked.GetSpan("0ca7ff81-5ffd-11eb-ac71-6c626debff99")
	span1.Finish()
	t.Log((span1))
}

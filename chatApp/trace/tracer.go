package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface for tracing events through code.
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}
type nilTracer struct{}

// New creates a new Tracer that writes to the provided io.Writer.
func New(w io.Writer) Tracer {
	return &tracer{out: w}
	//trace output 
}

func (t *tracer) Trace(a ...interface{}) {
	
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

func (t* nilTracer)Trace(a ...interface{}){}
  func Off() Tracer{
	return &nilTracer{}
  }
  //niltracer does what ?
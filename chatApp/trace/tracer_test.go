package trace

import (
	"bytes"
	"testing"
)
func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
	t.Error("Return from New should not be nil")
	} else {
	tracer.Trace("Hello trace package.")
	if buf.String() != "Hello trace package.\n" {
	t.Errorf("Trace should not write '%s'.", buf.String())
	}
	}
	}
	//red green testing ?
	//why do we write go client.Write instead of writing without go ?
	// channels infinte waiting for msg in go ?

func TestOff(t *testing.T){
	var silentTracer Tracer = Off()
	silentTracer.Trace("Hello trace package. time to go off")
}	
